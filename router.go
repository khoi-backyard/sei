package sei

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) Add(method, path string, h HandlerFunc) {
	r.handlers[path] = h
}

func (r *Router) Find(method, path string) HandlerFunc {
	return r.handlers[path]
}
