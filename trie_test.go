package sei

import (
	"testing"
)

var validRoutes = []struct {
	route string
}{
	{"/api/v1/"},
	{"/api/v1/products"},
	{"/api/v1/products/1"},
}

var invalidRoutes = []struct {
	route string
}{
	{"/api/v2/"},
	{"/api/v1/products/2"},
}

func TestTrie_Add(t *testing.T) {
	trie := NewTrie()

	data := "some associated data"
	n := trie.Add("/api/v1", data)

	if n.Data().(string) != data {
		t.Errorf("Expected %s, got: %s", data, n.Data().(string))
	}
}

func TestTrie_Find(t *testing.T) {
	trie := NewTrie()

	for _, c := range validRoutes {
		trie.Add(c.route, c.route)
	}

	for _, c := range validRoutes {
		n, ok := trie.Find(c.route)

		if !ok {
			t.Errorf("Expected to find %s", c.route)
		}

		if n.Data().(string) != c.route {
			t.Errorf("Expected %s, got: %s", c.route, n.Data().(string))
		}
	}

	for _, c := range invalidRoutes {
		if _, ok := trie.Find(c.route); ok {
			t.Errorf("Shouldn't found %s", c.route)
		}
	}
}
