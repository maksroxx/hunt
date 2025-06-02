package main

import (
	"github.com/maksroxx/hunt/hunt"
)

func main() {
	h := hunt.New()
	h.Debug(true)

	api := h.Group("/api")
	api.GET("/hello", func(c *hunt.Context) {
		c.String(200, "Hello from /api/hello")
	})
	h.Run(":8080")
}
