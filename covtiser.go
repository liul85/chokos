package covtiser

import (
	"fmt"
	"net/http"
)

type Engine struct {
	routes map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{make(map[string]http.HandlerFunc)}
}

func (engine *Engine) addRoute(method string, path string, handler http.HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, path)
	engine.routes[key] = handler
}

func (engine *Engine) Get(path string, handler http.HandlerFunc) {
	engine.addRoute(http.MethodGet, path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := fmt.Sprintf("%s-%s", r.Method, r.URL.Path)
	if handler, ok := engine.routes[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
	}
}

func (engine *Engine) Run(port string) (err error) {
	fmt.Printf("Server started on port %s!\n", port)
	return http.ListenAndServe(port, engine)
}
