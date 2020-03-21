package covtiser

import (
	"fmt"
	"net/http"
)

// Context is the encapsulation of http request
type Context struct {
	writer  http.ResponseWriter
	request *http.Request
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{w, req}
}

func (context *Context) String(statusCode int, template string, params ...interface{}) {
	context.writer.WriteHeader(statusCode)
	context.writer.Write([]byte(fmt.Sprintf(template, params...)))
}
