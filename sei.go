package sei

import (
	"net/http"
	"sync"
)

type (
	Sei struct {
		middlewares        []MiddlewareFunc
		router             *Router
		contextPool        sync.Pool
		invalidPathHandler HandlerFunc
	}
	MiddlewareFunc func(HandlerFunc) HandlerFunc
	HandlerFunc    func(*Context) error
)

func New() *Sei {
	sei := Sei{}
	sei.contextPool.New = func() interface{} {
		return NewContext()
	}

	sei.invalidPathHandler = func(c *Context) error {
		return c.String(http.StatusNotFound, "404")
	}

	sei.router = NewRouter()
	return &sei
}

func (s *Sei) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := s.contextPool.Get().(*Context)
	defer s.contextPool.Put(c)

	c.Reset(w, r)

	method := r.Method
	path := r.URL.RawPath
	if path == "" {
		path = r.URL.Path
	}

	h := s.router.Find(method, path)

	if h == nil {
		h = s.invalidPathHandler
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
	s.router.Add(http.MethodGet, path, h)
}

func (s *Sei) Use(mws ...MiddlewareFunc) {
	s.middlewares = append(s.middlewares, mws...)
}
