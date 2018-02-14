package trie

import "testing"

func TestTrie_Add(t *testing.T) {
	trie := New()

	key := "/abc"

	n := trie.Add(key, 123)

	found, ok := trie.Find(key)

	if !ok {
		t.Fatal("Can't find %s", key)
	}

	if found.Data().(int) != 123 {
		t.Errorf("Expected 123, got: %d", n.Data().(int))
	}
}
