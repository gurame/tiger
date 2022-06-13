package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/api"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome tiger api!"))
	})
	api.Map(app)
	log.Fatal(app.Listen(":3000"))
}
