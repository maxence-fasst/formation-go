package requests

import (
	"formation-go/helpers"

	"github.com/go-playground/validator/v10"
)

func NewCustomValidators() *helpers.CustomValidator {
	validate := validator.New()
	_ = validate.RegisterValidation("isValidSiren", isValidSiren)

	return &helpers.CustomValidator{Validator: validate}
}

func isValidSiren(fl validator.FieldLevel) bool {
	// TODO: Implement
	// value := fl.Field().String()
	return true
}
