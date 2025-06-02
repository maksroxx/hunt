package core

import (
	"strings"
)

type HandlerFunc func(*Context)

type Router struct {
	handlers map[string]HandlerFunc
	routes   map[string][]string
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
		routes:   make(map[string][]string),
	}
}

func (r *Router) AddRoute(method, pattern string, handlers ...HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = combineHandlers(handlers)
	r.routes[key] = strings.Split(pattern, "/")
}

func combineHandlers(handlers []HandlerFunc) HandlerFunc {
	return func(c *Context) {
		for _, h := range handlers {
			h(c)
		}
	}
}

func (r *Router) Handle(c *Context) {
	requestPath := strings.Split(c.Request.URL.Path, "/")
	method := c.Request.Method

	for key, segments := range r.routes {
		parts := strings.Split(key, "-")
		routeMethod := parts[0]

		if method != routeMethod {
			continue
		}

		if len(segments) != len(requestPath) {
			continue
		}

		params := make(map[string]string)
		matched := true

		for i := 0; i < len(segments); i++ {
			if strings.HasPrefix(segments[i], ":") {
				paramName := segments[i][1:]
				params[paramName] = requestPath[i]
			} else if segments[i] != requestPath[i] {
				matched = false
				break
			}
		}

		if matched {
			c.Params = params
			if handler, ok := r.handlers[key]; ok {
				c.DebugInfo()
				handler(c)
				return
			}
		}
	}

	c.String(404, "404 NOT FOUND: "+c.Request.URL.Path)
}
