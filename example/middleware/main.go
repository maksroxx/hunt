package main

import (
	"fmt"

	"github.com/maksroxx/hunt/hunt"
)

func Logger(c *hunt.Context) {
	fmt.Println("Logger start:", c.Request.URL.Path)
	c.Next()
	fmt.Println("Logger end:", c.Request.URL.Path)
}

func Auth(c *hunt.Context) {
	if c.Request.Header.Get("User") == "Roxx" {
		fmt.Println("Header Roxx")
	} else {
		fmt.Println("WTF")
	}
	c.Next()
}

func main() {
	h := hunt.New()
	h.Debug(true)

	api := h.Group("/api")
	api.Use(Logger, Auth)

	api.GET("/mid/:id", func(c *hunt.Context) {
		id := c.Param("id")
		c.JSON(200, map[string]string{"id": id})
	})
	h.Run(":8080")
}
