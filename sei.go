package sei

import (
	"net/http"
	"sync"
)

type (
	Sei struct {
		middlewares     []MiddlewareFunc
		router          *Router
		contextPool     sync.Pool
		notFoundHandler HandlerFunc
	}
	MiddlewareFunc func(HandlerFunc) HandlerFunc
	HandlerFunc    func(*Context) error
)

func New() *Sei {
	sei := Sei{}
	sei.contextPool.New = func() interface{} {
		return NewContext()
	}

	sei.notFoundHandler = func(c *Context) error {
		return c.String(http.StatusNotFound, "404")
	}

	sei.router = NewRouter(&sei)
	return &sei
}

func (s *Sei) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := s.getContext(w, r)
	defer s.putContext(c)

	method := r.Method
	path := r.URL.RawPath
	if path == "" {
		path = r.URL.Path
	}

	h := s.router.Find(method, path)

	if h == nil {
		h = s.notFoundHandler
	}

	for i := len(s.middlewares) - 1; i >= 0; i-- {
		h = s.middlewares[i](h)
	}

	h(c)
}

func (s *Sei) Server(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: s,
	}
}

func (s *Sei) Start(addr string) error {
	return s.Server(addr).ListenAndServe()
}

func (s *Sei) GET(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodGet)
}

func (s *Sei) HEAD(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodHead)
}

func (s *Sei) POST(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodPost)
}

func (s *Sei) PUT(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodPut)
}

func (s *Sei) PATCH(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodPatch)
}

func (s *Sei) DELETE(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodDelete)
}

func (s *Sei) CONNECT(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodConnect)
}

func (s *Sei) OPTIONS(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodOptions)
}

func (s *Sei) TRACE(path string, h HandlerFunc) {
	s.router.Add(h, path, http.MethodTrace)
}

func (s *Sei) Use(mws ...MiddlewareFunc) {
	s.middlewares = append(s.middlewares, mws...)
}

func (s *Sei) getContext(w http.ResponseWriter, r *http.Request) *Context {
	c := s.contextPool.Get().(*Context)
	c.Reset(w, r)
	return c
}

func (s *Sei) putContext(c *Context) {
	s.contextPool.Put(c)
}
