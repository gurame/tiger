package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/api/routes"
)

func Map(app fiber.Router) {
	routes.SellerRoute(app)
}
