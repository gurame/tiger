package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/app/infrastructure/db"
)

func List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sellers, err := db.Sellers().AllG(context.Background())
		if err != nil {
			log.Fatal("Error when listing sellers $s", err)
		}

		fmt.Println(sellers)
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   sellers,
			"error":  nil,
		})
	}
}
