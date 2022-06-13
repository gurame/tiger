package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/core/sellers/handlers"
)

func SellerRoute(app fiber.Router) {
	app.Post("/sellers", handlers.Add())
	app.Get("/sellers", handlers.Get())
}
