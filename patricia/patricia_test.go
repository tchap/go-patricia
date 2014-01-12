// Copyright (c) 2014 The go-patricia AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package patricia

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

const (
	success = true
	failure = false
)

type testData struct {
	key    string
	value  interface{}
	retVal bool
}

func TestTrie_InsertDifferentPrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_InsertDuplicatePrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepan", "Pepan Zdepan", failure},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_InsertVariousPrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepin", "Pepin Omacka", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Pepan", "Pepan Dupan", failure},
		{"Karel", "Karel Pekar", success},
		{"Jenik", "Jenik Poustevnicek", failure},
		{"Pepanek", "Pepanek Zemlicka", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_Visit(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	if err := trie.Visit(func(prefix Prefix, item Item) error {
		name := data[item.(int)].key
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if !strings.HasPrefix(string(prefix), name) {
			t.Errorf("Unexpected prefix encountered, %q not a prefix of %q", prefix, name)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestTrie_VisitSkipSubtree(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	if err := trie.Visit(func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if item.(int) == 0 {
			t.Logf("SKIP %q", prefix)
			return SkipSubtree
		}
		if strings.HasPrefix(string(prefix), "Pepa") {
			t.Errorf("Unexpected prefix encountered, %q", prefix)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestTrie_VisitReturnError(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	someErr := errors.New("Something exploded")
	if err := trie.Visit(func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if item.(int) == 0 {
			return someErr
		}
		if item.(int) != 0 {
			t.Errorf("Unexpected prefix encountered, %q", prefix)
		}
		return nil
	}); err != nil && err != someErr {
		t.Fatal(err)
	}
}

func TestTrie_VisitSubtree(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	var counter int
	subtreePrefix := []byte("Pep")
	if err := trie.VisitSubtree(subtreePrefix, func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if !bytes.HasPrefix(prefix, subtreePrefix) {
			t.Errorf("Unexpected prefix encountered, %q does not extend %q",
				prefix, subtreePrefix)
		}
		counter++
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if counter != 3 {
		t.Error("Unexpected number of nodes visited")
	}
}

func TestTrie_VisitPrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"P", 0, success},
		{"Pe", 1, success},
		{"Pep", 2, success},
		{"Pepa", 3, success},
		{"Pepa Zdepa", 4, success},
		{"Pepa Kuchar", 5, success},
		{"Honza", 6, success},
		{"Jenik", 7, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	var counter int
	word := []byte("Pepa")
	if err := trie.VisitPrefixes(word, func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if !bytes.HasPrefix(word, prefix) {
			t.Errorf("Unexpected prefix encountered, %q is not a prefix of %q",
				prefix, word)
		}
		counter++
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if counter != 4 {
		t.Error("Unexpected number of nodes visited")
	}
}

func TestParticiaTrie_Delete(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range data {
		t.Logf("DELETE word=%v, success=%v", v.key, v.retVal)
		if ok := trie.Delete([]byte(v.key)); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestParticiaTrie_DeleteNonExistent(t *testing.T) {
	trie := NewTrie()

	insertData := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
	}
	deleteData := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Pepan", "Pepan Zdepan", failure},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Honza", "Honza Novak", failure},
	}

	for _, v := range insertData {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range deleteData {
		t.Logf("DELETE word=%v, success=%v", v.key, v.retVal)
		if ok := trie.Delete([]byte(v.key)); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestParticiaTrie_DeleteSubtree(t *testing.T) {
	trie := NewTrie()

	insertData := []testData{
		{"P", 0, success},
		{"Pe", 1, success},
		{"Pep", 2, success},
		{"Pepa", 3, success},
		{"Pepa Zdepa", 4, success},
		{"Pepa Kuchar", 5, success},
		{"Honza", 6, success},
		{"Jenik", 7, success},
	}
	deleteData := []testData{
		{"Pe", -1, success},
		{"Pe", -1, failure},
		{"Honzik", -1, failure},
		{"Honza", -1, success},
		{"Honza", -1, failure},
		{"Pep", -1, failure},
		{"P", -1, success},
		{"Nobody", -1, failure},
		{"", -1, success},
	}

	for _, v := range insertData {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range deleteData {
		t.Logf("DELETE_SUBTREE prefix=%v, success=%v", v.key, v.retVal)
		if ok := trie.DeleteSubtree([]byte(v.key)); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

/*
func TestTrie_Dump(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Honda", nil, success},
		{"Honza", nil, success},
		{"Jenik", nil, success},
		{"Pepan", nil, success},
		{"Pepin", nil, success},
	}

	for i, v := range data {
		if _, ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Logf("INSERT %v %v", v.key, v.value)
			t.Fatalf("Unexpected return value, expected=%v, got=%v", i, ok)
		}
	}

	dump := `
+--+--+ Hon +--+--+ da
   |           |
   |           +--+ za
   |
   +--+ Jenik
   |
   +--+ Pep +--+--+ an
               |
               +--+ in
`

	var buf bytes.Buffer
	trie.Dump(buf)

	if !bytes.Equal(buf.Bytes(), dump) {
		t.Logf("DUMP")
		t.Fatalf("Unexpected dump generated, expected\n\n%v\ngot\n\n%v", dump, buf.String())
	}
}
*/

func TestTrie_longestCommonPrefixLenght(t *testing.T) {
	trie := NewTrie()
	trie.prefix = []byte("1234567890")

	switch {
	case trie.longestCommonPrefixLength([]byte("")) != 0:
		t.Fail()
	case trie.longestCommonPrefixLength([]byte("12345")) != 5:
		t.Fail()
	case trie.longestCommonPrefixLength([]byte("123789")) != 3:
		t.Fail()
	case trie.longestCommonPrefixLength([]byte("12345678901")) != 10:
		t.Fail()
	}
}
