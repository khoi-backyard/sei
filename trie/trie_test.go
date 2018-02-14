package trie

import "testing"

var tests = []struct {
	route string
}{
	{"/api/v1/"},
	{"/api/v1/products"},
	{"/api/v1/products/1"},
}

func TestTrie_Add(t *testing.T) {
	trie := New()

	data := "some associated data"
	n := trie.Add("/api/v1", data)

	if n.Data().(string) != data {
		t.Errorf("Expected %s, got: %s", data, n.Data().(string))
	}
}

func TestTrie_Find(t *testing.T) {
	trie := New()
	for _, c := range tests {
		trie.Add(c.route, c.route)
	}

	for _, c := range tests {
		n, ok := trie.Find(c.route)

		if !ok {
			t.Errorf("Expected to find %s", c.route)
		}

		if n.Data().(string) != c.route {
			t.Errorf("Expected %s, got: %s", c.route, n.Data().(string))
		}
	}
}
