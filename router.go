package chokos

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{roots: make(map[string]*node), handlers: make(map[string]HandlerFunc)}
}

func (router *router) addRoute(method string, path string, handler HandlerFunc) {
	parts := parsePattern(path)

	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}

	key := fmt.Sprintf("%s-%s", method, path)
	router.roots[method].insert(path, parts, 0)
	router.handlers[key] = handler
}

func (router *router) handle(context *Context) {
	node, params := router.getRoute(context.request.Method, context.request.URL.Path)
	if node != nil {
		context.params = params
		key := fmt.Sprintf("%s-%s", context.request.Method, node.pattern)
		context.handlers = append(context.handlers, router.handlers[key])
	} else {
		context.handlers = append(context.handlers, func(c *Context) {
			c.String(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s\n", context.request.URL.Path))
		})
	}

	context.Next()
}

func (router *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := router.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts, "/")
				break
			}
		}

		return n, params
	}

	return nil, nil
}

func parsePattern(pattern string) []string {
	patternSplits := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range patternSplits {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// Use is able to allow user to use middlewares
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
