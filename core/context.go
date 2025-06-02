package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Context struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	Params    map[string]string
	DebugMode bool
	handlers  []HandlerFunc
	index     int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Params:  make(map[string]string),
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) String(code int, message string) {
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(message))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) DebugInfo() {
	if !c.DebugMode {
		return
	}

	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Cyan   = "\033[36m"
		Bold   = "\033[1m"
	)

	fmt.Println(Bold + Cyan + "┌─────[HUNT DEBUG]──────" + Reset)
	fmt.Printf(Bold+"│ Method: "+Reset+"%s\n", c.Request.Method)
	fmt.Printf(Bold+"│ URL:    "+Reset+"%s\n", c.Request.URL.String())
	fmt.Println(Bold + "│ Headers:" + Reset)
	for k, v := range c.Request.Header {
		fmt.Printf("│   "+Yellow+"%s"+Reset+": %v\n", k, v)
	}

	if c.Request.Body != nil {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err == nil && len(bodyBytes) > 0 {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			fmt.Println(Bold + "│ Body:" + Reset)
			bodyLines := strings.Split(string(bodyBytes), "\n")
			for _, line := range bodyLines {
				fmt.Printf("│   %s\n", line)
			}
		}
	}

	fmt.Println(Bold + Cyan + "└───────────────────────" + Reset)
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) SetHandlers(h []HandlerFunc) {
	c.handlers = h
	c.index = -1
}
