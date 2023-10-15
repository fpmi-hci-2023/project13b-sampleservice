package main

import "github.com/gofiber/fiber/v2"

var (
	Version string
	Commit  string
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!\nServer version: " + Version + ", commit: " + Commit + "\n")
	})

	app.Listen(":8080")
}
