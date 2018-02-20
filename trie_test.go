package sei

import "testing"

func TestTrie_Find(t *testing.T) {
	s := New()

	s.GET("/products/:id/colors/:cid", func(ctx *Context) error {
		return nil
	})

	dumpTree(s.router.tree.root)
}
