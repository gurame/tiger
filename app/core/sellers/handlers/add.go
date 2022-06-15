package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dbcontext "github.com/gurame/tiger/app/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type (
	SellerCreateReq struct {
		TaxId string `json:"taxId" validate:"required"`
		Name  string `json:"name" validate:"required"`
	}

	ErrorResponse struct {
		FailedField string
		Tag         string
		Value       string
	}
)

var validate = validator.New()

func ValidateStruct(seller SellerCreateReq) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(seller)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// Add registers a new seller
// @Summary Register a new seller
// @Description Register seller
// @Tags sellers
// @Accept json
// @Produce json
// @Param seller body handlers.SellerCreateReq true "Register seller"
// @Success 200 {object} handlers.SellerCreateReq
// @Failure 400 {object} ErrorResponse{}
// @Router /sellers [post]
func Add() fiber.Handler {
	return func(c *fiber.Ctx) error {
		r := new(SellerCreateReq)
		if err := c.BodyParser(r); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		errors := ValidateStruct(*r)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		s := &dbcontext.Seller{Sellerid: uuid.NewString(), Name: r.Name, Taxid: r.TaxId}

		s.InsertG(context.Background(), boil.Infer())

		return c.Status(fiber.StatusCreated).JSON(r)
	}
}
