package covtiser

import (
	"fmt"
	"net/http"
)

// Engine  represents HTTP request process unit
type Engine struct {
	router *router
}

// New function creates a new Engine
func New() *Engine {
	return &Engine{newRouter()}
}

// Get method add a path mapping for HTTP Get method
func (engine *Engine) Get(path string, handler http.HandlerFunc) {
	engine.router.addRoute(http.MethodGet, path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	engine.router.handle(w, r)
}

// Run method starts the app
func (engine *Engine) Run(port string) (err error) {
	fmt.Printf("Server started on port %s!\n", port)
	return http.ListenAndServe(port, engine)
}
