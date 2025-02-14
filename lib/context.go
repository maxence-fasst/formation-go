package lib

import (
	"formation-go/helpers"

	"github.com/labstack/echo/v4"
)

func BindAndValidate[T any](context echo.Context) (*T, error) {
	inputContext := new(T)

	var err error

	if err = context.Bind(inputContext); err != nil {
		return nil, err
	}

	helpers.PrintData("Request with data ", &inputContext)

	if err = context.Validate(inputContext); err != nil {
		return nil, err
	}

	return inputContext, nil
}
