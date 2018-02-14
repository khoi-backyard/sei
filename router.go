package sei

import (
	"net/http"

	"github.com/armon/go-radix"
)

type Router struct {
	tree *radix.Tree
	sei  *Sei
}

func NewRouter(s *Sei) *Router {
	return &Router{
		sei:  s,
		tree: radix.New(),
	}
}

func (r *Router) Add(method, path string, h HandlerFunc) {
	r.tree.Insert(path, h)
}

func (r *Router) Find(method, path string) HandlerFunc {
	_, h, _ := r.tree.LongestPrefix(path)
	return h.(HandlerFunc)
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
