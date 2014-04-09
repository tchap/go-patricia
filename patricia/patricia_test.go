// Copyright (c) 2014 The go-patricia AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package patricia

import (
	"testing"
)

// Tests -----------------------------------------------------------------------

func TestTrie_GetNonexistentPrefix(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"aba", 0, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	t.Logf("GET prefix=baa, expect item=nil")
	if item := trie.Get(Prefix("baa")); item != nil {
		t.Errorf("Unexpected return value, expected=<nil>, got=%v", item)
	}
}
