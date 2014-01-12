// Copyright (c) 2014 The go-patricia AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package patricia

import "errors"

type (
	Prefix      []byte
	Item        interface{}
	VisitorFunc func(prefix Prefix, item Item) error
)

// Trie is a generic patricia trie that allows
//
// 1. visit all items saved in the tree,
// 2. visit all items matching particular prefix (visit subtree), or
// 3. given a string, visit all items matching some prefix of the string.
//
// Trie is not thread-safe.
type Trie struct {
	prefix Prefix
	item   Item

	children childList
}

// Trie constructor.
func NewTrie() *Trie {
	return &Trie{
		children: newSparseChildList(),
	}
}

// Item returns the item stored in the root of this trie.
func (trie *Trie) Item() Item {
	return trie.item
}

// Insert inserts a new item into the trie using the given prefix.
//
// True is returned unless the node with the given key already exists.
func (trie *Trie) Insert(prefix Prefix, item Item) (inserted bool) {
	// Nil prefix not allowed.
	if prefix == nil {
		panic(ErrNilPrefix)
	}

	var (
		common int
		node   *Trie = trie
		child  *Trie
	)

	if trie.prefix == nil {
		goto ExtendPath
	}

	for {
		// Compute the longest common prefix length.
		common = node.longestCommonPrefixLength(prefix)
		prefix = prefix[common:]

		// Only a part matches, split.
		if common < len(node.prefix) {
			goto SplitPrefix
		}

		// common == len(node.prefix) since never (common > len(node.prefix))
		// common == len(former prefix) <-> 0 == len(prefix)
		// -> former prefix == node.prefix
		if len(prefix) == 0 {
			goto InsertItem
		}

		// Check children for matching prefix.
		child = node.children.next(prefix[0])
		if child == nil {
			goto AppendChild
		}
		node = child
	}

SplitPrefix:
	// Split the prefix if necessary.
	child = new(Trie)
	*child = *node
	*node = *NewTrie()
	node.prefix = child.prefix[:common]
	child.prefix = child.prefix[common:]
	child = child.compact()
	node.children = node.children.add(child)

AppendChild:
	child = NewTrie()
	node.children = node.children.add(child)
	node = child

ExtendPath:
	// Keep appending children until whole prefix is inserted.
	for {
		if len(prefix) <= MaxPrefixPerNode {
			node.prefix = prefix
			goto InsertItem
		} else {
			child := NewTrie()
			child.prefix = node.prefix[MaxPrefixPerNode:]
			node.prefix = node.prefix[:MaxPrefixPerNode]
			node.children = node.children.add(child)
			node = child
		}
	}

InsertItem:
	// Try to insert the item if possible.
	if node.item == nil {
		node.item = item
		inserted = true
		return
	}
	return
}

// Visit calls visitor on every node containing a non-nil item.
//
// If an error is returned from visitor, the function stops visiting the tree
// and returns that error, unless it is a special error - SkipSubtree. In that
// case Visit skips the subtree represented by the current node and continues
// elsewhere.
func (trie *Trie) Visit(visitor VisitorFunc) error {
	return trie.walk(nil, visitor)
}

// VisitSubtree works much like Visit, but it only visits nodes matching prefix.
func (trie *Trie) VisitSubtree(prefix Prefix, visitor VisitorFunc) error {
	// Nil prefix not allowed.
	if prefix == nil {
		panic(ErrNilPrefix)
	}

	// Empty trie must be handled explicitly.
	if trie.prefix == nil {
		return nil
	}

	// Locate the relevant subtree.
	_, root, found, _ := trie.findSubtree(prefix)
	if !found {
		return nil
	}

	// Visit it.
	return root.walk(prefix, visitor)
}

// VisitPrefixes visits only nodes that represent prefixes of word.
// To say the obvious, returning SkipSubtree from visitor makes no sense here.
func (trie *Trie) VisitPrefixes(word Prefix, visitor VisitorFunc) error {
	// Nil word not allowed.
	if word == nil {
		panic(ErrNilPrefix)
	}

	// Empty trie must be handled explicitly.
	if trie.prefix == nil {
		return nil
	}

	// Walk the path matching word prefixes.
	node := trie
	prefix := word
	offset := 0
	for {
		// Compute what part of prefix matches.
		common := node.longestCommonPrefixLength(word)
		word = word[common:]
		offset += common

		// Partial match means that there is no subtree matching prefix.
		if common < len(node.prefix) {
			return nil
		}

		// Call the visitor.
		if item := node.item; item != nil {
			if err := visitor(prefix[:offset], item); err != nil {
				return err
			}
		}

		if len(word) == 0 {
			// This node represents word, we are finished.
			return nil
		}

		// There is some word suffix left, move to the children.
		child := node.children.next(word[0])
		if child == nil {
			// There is nowhere to continue, return.
			return nil
		}

		node = child
	}
}

// Delete deletes the item represented by the given prefix.
//
// True is returned if the matching node was found and deleted.
func (trie *Trie) Delete(word Prefix) (deleted bool) {
	// Nil prefix not allowed.
	if word == nil {
		panic(ErrNilPrefix)
	}

	// Empty trie must be handled explicitly.
	if trie.prefix == nil {
		return false
	}

	// Find the relevant node.
	parent, node, _, match := trie.findSubtree(word)
	if !match {
		return false
	}

	// If the item is already set to nil, there is nothing to do.
	if node.item == nil {
		return false
	}

	// Delete the item.
	node.item = nil

	// Compact since that might be possible now.
	node = node.compact()
	if parent != nil {
		*parent = *parent.compact()
	}
	parent.children.replace(node)

	return true
}

// DeleteSubtree finds the subtree exactly matching prefix and deletes it.
//
// True is returned if the subtree was found and deleted.
func (trie *Trie) DeleteSubtree(prefix Prefix) (deleted bool) {
	// Nil prefix not allowed.
	if prefix == nil {
		panic(ErrNilPrefix)
	}

	// Empty trie must be handled explicitly.
	if trie.prefix == nil {
		return false
	}

	// Locate the relevant subtree.
	parent, root, found, _ := trie.findSubtree(prefix)
	if !found {
		return false
	}

	// If we are in the root of the trie, reset the trie.
	if parent == nil {
		root.prefix = nil
		root.children = newSparseChildList()
		return true
	}

	// Otherwise remove the root node from its parent.
	parent.children.remove(root)
	return true
}

func (trie *Trie) findSubtree(prefix Prefix) (parent *Trie, root *Trie, found bool, exactMatch bool) {
	// Find the subtree matching prefix.
	root = trie
	for {
		// Compute what part of prefix matches.
		common := root.longestCommonPrefixLength(prefix)
		prefix = prefix[common:]

		// We used up the whole prefix, subtree found.
		if len(prefix) == 0 {
			found = true
			if common == len(root.prefix) {
				exactMatch = true
			}
			return
		}

		// Partial match means that there is no subtree matching prefix.
		if common < len(root.prefix) {
			return
		}

		// There is some prefix left, move to the children.
		child := root.children.next(prefix[0])
		if child == nil {
			// There is nowhere to continue, there is no subtree matching prefix.
			return
		}

		parent = root
		root = child
	}
}

func (trie *Trie) walk(actualRootPrefix Prefix, visitor VisitorFunc) error {
	var prefix Prefix
	// Allocate a bit more space for prefix at the beginning.
	if actualRootPrefix == nil {
		prefix = make(Prefix, 32)
		copy(prefix, trie.prefix)
		prefix = prefix[:len(trie.prefix)]
	} else {
		prefix = make(Prefix, 32+len(actualRootPrefix))
		copy(prefix, actualRootPrefix)
		copy(prefix[len(actualRootPrefix):], trie.prefix)
		prefix = prefix[:len(actualRootPrefix)+len(trie.prefix)]
	}

	// Visit the root first. Not that this works for empty trie as well since
	// in that case item == nil && len(children) == 0.
	if trie.item != nil {
		if err := visitor(prefix, trie.item); err != nil {
			if err == SkipSubtree {
				return nil
			}
			return err
		}
	}

	// Then continue to the children.
	return trie.children.walk(&prefix, visitor)
}

func (trie *Trie) compact() *Trie {
	// Only a node with a single child can be compacted.
	if trie.children.length() != 1 {
		return trie
	}

	child := trie.children.head()

	// Only nodes that are not both containing items can be compacted.
	if trie.item != nil && child.item != nil {
		return trie
	}

	// Make sure the combined prefixes fit into a single node.
	if len(trie.prefix)+len(child.prefix) > MaxPrefixPerNode {
		return trie
	}

	// Concatenate the prefixes, move the items.
	child.prefix = append(trie.prefix, child.prefix...)
	if trie.item != nil {
		child.item = trie.item
	}

	return child
}

func (trie *Trie) longestCommonPrefixLength(prefix Prefix) (i int) {
	for ; i < len(prefix) && i < len(trie.prefix) && prefix[i] == trie.prefix[i]; i++ {
	}
	return
}

// Errors ----------------------------------------------------------------------

var (
	SkipSubtree  = errors.New("Skip this subtree")
	ErrNilPrefix = errors.New("Nil prefix passed into a method call")
)
