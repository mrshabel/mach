package main

import (
	"log"

	"github.com/mrshabel/mach"
)

func main() {
	// create default instance
	app := mach.Default()

	// plain text response
	app.GET("/", func(c *mach.Context) {
		c.Text(200, "Welcome to the basic tutorial")
	})

	// json response
	app.GET("/hello/{name}", func(c *mach.Context) {
		name := c.Param("name")

		c.JSON(200, map[string]any{
			"message": "Hello " + name,
		})
	})

	// run app
	log.Fatal(app.Run("localhost:8000"))
}
