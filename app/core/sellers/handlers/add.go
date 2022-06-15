package handlers

import "github.com/gofiber/fiber/v2"

func Add() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": true})
	}
}
