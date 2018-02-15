package sei

import (
	"net/http"
)

type Router struct {
	tree *Trie
	sei  *Sei
}

func NewRouter(s *Sei) *Router {
	return &Router{
		sei:  s,
		tree: NewTrie(),
	}
}

func (r *Router) Add(method, path string, h HandlerFunc) {
	r.tree.Add(path, h)
}

func (r *Router) Find(method, path string) HandlerFunc {
	n, ok := r.tree.Find(path)

	if !ok {
		return nil
	}

	h, ok := n.Data().(HandlerFunc)

	if !ok {
		return nil
	}

	return h
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := r.sei.getContext(w, req)
	defer r.sei.putContext(c)

	h := r.Find(req.Method, req.URL.Path)

	if h == nil {
		h = r.sei.notFoundHandler
	}

	h(c)
}
