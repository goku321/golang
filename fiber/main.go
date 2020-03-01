package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Use(first, second)
	app.Use("/api", third)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Api Router Handler")
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

var third = func(c *fiber.Ctx) {
	fmt.Println("Third Middleware")
	c.Next()
}
