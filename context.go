package chokos

import (
	"encoding/json"
	"net/http"
)

// Context is the encapsulation of http request
type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	params  map[string]string
}

// Param method return value for specified parameter
func (context *Context) Param(key string) string {
	value, _ := context.params[key]
	return value
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{w, req, make(map[string]string)}
}

func (context *Context) String(statusCode int, content string) {
	context.writer.Header().Set("Content-Type", "text/plain")
	context.writer.WriteHeader(statusCode)
	context.writer.Write([]byte(content))
}

// JSON method allows user to return json response in response
func (context *Context) JSON(statusCode int, obj interface{}) {
	context.writer.Header().Set("Content-Type", "application/json")
	context.writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(context.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.writer, err.Error(), 500)
	}
}

// HTML method allows user to return html in response
func (context *Context) HTML(statusCode int, html string) {
	context.writer.Header().Set("Content-Type", "text/html")
	context.writer.WriteHeader(statusCode)
	context.writer.Write([]byte(html))
}

// Data method allows user to send data in response
func (context *Context) Data(statusCode int, data []byte) {
	context.writer.WriteHeader(statusCode)
	context.writer.Write(data)
}
