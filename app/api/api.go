package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gurame/tiger/app/api/routes"
	"github.com/gurame/tiger/app/infrastructure/config"
	_ "github.com/lib/pq"
)

func Run() {

	//db
	config.Database()

	//web
	app := fiber.New()
	app.Use(recover.New())

	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome tiger api!"))
	})
	routes.ConfigureSeller(app)

	log.Fatal(app.Listen(":3000"))
}
