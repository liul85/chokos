package chokos

import (
	"fmt"
	"net/http"
)

// HandlerFunc is the func that handle each route mapping
type HandlerFunc func(*Context)

// Engine  represents HTTP request process unit
type Engine struct {
	router *router
}

// New function creates a new Engine
func New() *Engine {
	return &Engine{newRouter()}
}

// Get method add a path mapping for HTTP Get method
func (engine *Engine) Get(path string, handler HandlerFunc) {
	engine.router.addRoute(http.MethodGet, path, handler)
}

// Post method add a path mapping for HTTP Post method
func (engine *Engine) Post(path string, handler HandlerFunc) {
	engine.router.addRoute(http.MethodPost, path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	engine.router.handle(context)
}

// Run method starts the app
func (engine *Engine) Run(port string) (err error) {
	fmt.Printf("Server started on port %s!\n", port)
	return http.ListenAndServe(port, engine)
}
