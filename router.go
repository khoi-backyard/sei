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

func (r *Router) Add(h HandlerFunc, path string, methods ...string) {
	r.tree.Add(h, path, methods...)
}

func (r *Router) Find(method, path string) HandlerFunc {
	return r.tree.Find(path, method)
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
