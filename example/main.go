package main

import "github.com/maksroxx/hunt/hunt"

func main() {
	r := hunt.New()
	r.Debug(true)

	r.GET("/", func(c *hunt.Context) {
		c.String(200, "Welcome to Hunt!")
	})

	r.GET("/hello", func(c *hunt.Context) {
		name := c.Query("name")
		if name == "" {
			name = "stranger"
		}
		c.String(200, "Hello "+name)
	})

	r.GET("/users/:id", func(c *hunt.Context) {
		id := c.Param("id")
		c.JSON(200, map[string]string{"user_id": id})
	})

	r.DELETE("/users/:id", func(c *hunt.Context) {
		id := c.Param("id")
		c.String(200, "Deleted user "+id)
	})

	r.Run(":8080")
}
