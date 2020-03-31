package chokos

import (
	"fmt"
	"net/http"
	"strings"
)

// HandlerFunc is the func that handle each route mapping
type HandlerFunc func(*Context)

// Engine  represents HTTP request process unit
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New function creates a new Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//Group method is used to add new group to engine
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
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
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := newContext(w, r)
	context.handlers = middlewares
	engine.router.handle(context)
}

// Run method starts the app
func (engine *Engine) Run(port string) (err error) {
	fmt.Printf("Server started on port %s!\n", port)
	return http.ListenAndServe(port, engine)
}

// RouterGroup router group for access control
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (group *RouterGroup) addRoute(method string, path string, handler HandlerFunc) {
	pattern := group.prefix + path
	group.engine.router.addRoute(method, pattern, handler)
}

//Get allows user to create route for Get method
func (group *RouterGroup) Get(path string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, path, handler)
}

//Post allows user to create route for Post method
func (group *RouterGroup) Post(path string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, path, handler)
}
