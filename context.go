package sei

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	req  *http.Request
	res  *Response
	data map[string]interface{}
}

func NewContext() *Context {
	return &Context{
		data: make(map[string]interface{}),
		res:  new(Response),
	}
}

func (c *Context) Request() *http.Request {
	return c.req
}

func (c *Context) Response() *Response {
	return c.res
}

func (c *Context) String(statusCode int, s string) error {
	c.res.Writer.Header().Set("Content-Type", "text/plain")
	c.res.Writer.WriteHeader(statusCode)
	c.res.Writer.Write([]byte(s))
	return nil
}

func (c *Context) JSON(statusCode int, payload interface{}) error {
	c.res.Writer.Header().Set("Content-Type", "application/json")
	c.res.Writer.WriteHeader(statusCode)
	return json.NewEncoder(c.res.Writer).Encode(payload)
}

func (c *Context) Redirect(statusCode int, url string) {
	http.Redirect(c.res.Writer, c.req, url, statusCode)
}

func (c *Context) Set(key string, val interface{}) {
	c.data[key] = val
}

func (c *Context) Get(key string) interface{} {
	return c.data[key]
}

func (c *Context) Reset(w http.ResponseWriter, r *http.Request) {
	c.req = r
	c.res.Writer = w
}
