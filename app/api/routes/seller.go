package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/core/sellers/handlers"
)

func ConfigureSeller(app fiber.Router) {
	app.Get("/sellers", handlers.List())
	app.Get("/sellers/:id", handlers.Get())
	app.Post("/sellers", handlers.Add())
	app.Put("/sellers/:id", handlers.Update())
	app.Delete("/sellers/:id", handlers.Delete())
}
