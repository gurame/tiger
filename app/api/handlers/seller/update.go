package handlers

import "github.com/gofiber/fiber/v2"

func Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{"name": "update", "id": id})
	}
}
