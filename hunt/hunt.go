package hunt

import (
	"fmt"
	"net/http"

	"github.com/maksroxx/hunt/core"
)

type Context = core.Context
type HandlerFunc = core.HandlerFunc

type Engine struct {
	router      *core.Router
	debugMode   bool
	middlewares []HandlerFunc
}

type RouterGroup struct {
	prefix      string
	engine      *Engine
	middlewares []HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: core.NewRouter(),
	}
}

func (e *Engine) Debug(enabled bool) {
	e.debugMode = enabled
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := core.NewContext(w, req)
	c.DebugMode = e.debugMode
	e.router.Handle(c)
}

func (e *Engine) Run(addr string) error {
	printBanner()
	fmt.Printf("Hunt server is running at %s\n", addr)
	return http.ListenAndServe(addr, e)
}

func printBanner() {
	banner := `
  ░▀█░█████████████████▀▀░░░██░████
  ▄▄█████████████████▀░░░░░░██░████
  ███▀▀████████████▀░░░░░░░▄█░░████
  ███▄░░░░▀▀█████▀░▄▀▄░░░░▄█░░▄████
  ░███▄▄░░▄▀▄░▀███▄▀▀░░▄▄▀█▀░░█████
  ▄▄█▄▀█▄▄░▀▀████████▀███░░▄░██████
  ▀████▄▀▀▀██▀▀██▀▀██░░▀█░░█▄█████░
  ░░▀▀███▄░▀█░░▀█░░░▀░█░░░▄██████░▄
  ████▄▄▀██▄▄▄░█▄▄░▄█▄█▄███████░░░█
`
	fmt.Println(banner)
}

// Base route registration
func (e *Engine) addRoute(method, path string, handlers ...HandlerFunc) {
	e.router.AddRoute(method, path, handlers...)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.addRoute("PUT", path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.addRoute("DELETE", path, handler)
}

func (e *Engine) PATCH(path string, handler HandlerFunc) {
	e.addRoute("PATCH", path, handler)
}

// Router Groups
func (e *Engine) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: prefix,
		engine: e,
	}
}

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	fullPath := g.prefix + comp
	finalHandlers := append(g.middlewares, handler)
	g.engine.router.AddRoute(method, fullPath, finalHandlers...)
}

func (g *RouterGroup) GET(path string, handler HandlerFunc) {
	g.addRoute("GET", path, handler)
}

func (g *RouterGroup) POST(path string, handler HandlerFunc) {
	g.addRoute("POST", path, handler)
}

func (g *RouterGroup) PUT(path string, handler HandlerFunc) {
	g.addRoute("PUT", path, handler)
}

func (g *RouterGroup) DELETE(path string, handler HandlerFunc) {
	g.addRoute("DELETE", path, handler)
}

func (g *RouterGroup) PATCH(path string, handler HandlerFunc) {
	g.addRoute("PATCH", path, handler)
}
