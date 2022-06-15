package common

import (
	"github.com/go-playground/validator"
	"github.com/gurame/tiger/app/core/models"
)

func Validate(s interface{}) []*models.ErrorResponse {
	var validate = validator.New()
	var errors []*models.ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
