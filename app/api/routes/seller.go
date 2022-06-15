package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/gurame/tiger/app/api/handlers/seller"
	"github.com/gurame/tiger/app/core/services"
)

func ConfigureSeller(app fiber.Router, service services.ISellerService) {
	app.Get("/sellers", handlers.List())
	app.Get("/sellers/:id", handlers.Get())
	app.Post("/sellers", handlers.Add(service))
	app.Put("/sellers/:id", handlers.Update())
	app.Delete("/sellers/:id", handlers.Delete())
}
