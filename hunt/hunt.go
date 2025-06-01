package hunt

import (
	"fmt"
	"net/http"

	"github.com/maksroxx/hunt/core"
)

type Context = core.Context
type HandlerFunc = core.HandlerFunc

type Engine struct {
	router    *core.Router
	debugMode bool
}

func New() *Engine {
	return &Engine{
		router:    core.NewRouter(),
		debugMode: false,
	}
}

func (e *Engine) Debug(enabled bool) {
	e.debugMode = enabled
}

func (e *Engine) GET(path string, handler core.HandlerFunc) {
	e.router.AddRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler core.HandlerFunc) {
	e.router.AddRoute("POST", path, handler)
}

func (e *Engine) DELETE(path string, hanler core.HandlerFunc) {
	e.router.AddRoute("DELETE", path, hanler)
}

func (e *Engine) Run(addr string) error {
	printBanner()
	fmt.Printf("Hunt server is running at %s\n", addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := core.NewContext(w, req)
	c.DebugMode = e.debugMode
	e.router.Handle(c)
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
