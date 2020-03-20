package covtiser

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]http.HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]http.HandlerFunc)}
}

func (router *router) addRoute(method string, path string, handler http.HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, path)
	router.handlers[key] = handler
}

func (router *router) handle(w http.ResponseWriter, r *http.Request) {
	key := fmt.Sprintf("%s-%s", r.Method, r.URL.Path)
	if handler, ok := router.handlers[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
	}
}
