package chokos

import (
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

func (context *Context) String(statusCode int, content string) {
	context.writer.WriteHeader(statusCode)
	context.writer.Write([]byte(content))
}
