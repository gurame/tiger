package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/infrastructure/db"
)

func List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sellers, _ := db.Sellers().AllG(context.Background())

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   sellers,
			"error":  nil,
		})
	}
}
