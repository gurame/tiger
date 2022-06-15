package handlers

import "github.com/gofiber/fiber/v2"

func Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{"name": "delete", "id": id})
	}
}
