package covtiser

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (router *router) addRoute(method string, path string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, path)
	router.handlers[key] = handler
}

func (router *router) handle(context *Context) {
	key := fmt.Sprintf("%s-%s", context.request.Method, context.request.URL.Path)
	if handler, ok := router.handlers[key]; ok {
		handler(context)
	} else {
		context.String(http.StatusNotFound, "404 NOT FOUND: %s\n", context.request.URL.Path)
	}
}
