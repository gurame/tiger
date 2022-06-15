package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/app/core/models"
	"github.com/gurame/tiger/app/core/services"
)

// Add registers a new seller
// @Summary Register a new seller
// @Description Register seller
// @Tags sellers
// @Accept json
// @Produce json
// @Param seller body models.SellerCreateReq true "Register seller"
// @Success 200 {object} models.SellerCreateRes
// @Failure 400 {object} models.ErrorResponse{}
// @Router /sellers [post]
func Add(service services.ISellerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		r := new(models.SellerCreateReq)
		if err := c.BodyParser(r); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		res, _ := service.Add(*r)

		return c.Status(fiber.StatusCreated).JSON(res)
	}
}
