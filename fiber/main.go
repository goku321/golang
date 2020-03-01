package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Use(first, second)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(3000)
}

var first = func(c *fiber.Ctx) {
	fmt.Println("First Middleware")
	c.Next()
}

var second = func(c *fiber.Ctx) {
	fmt.Println("Second Middleware")
	c.Next()
}
