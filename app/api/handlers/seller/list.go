package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	dbcontext "github.com/gurame/tiger/app/infrastructure/db/models"
)

func List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sellers, err := dbcontext.Sellers().AllG(context.Background())
		if err != nil {
			log.Fatal("Error when listing sellers $s", err)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   sellers,
			"error":  nil,
		})
	}
}
