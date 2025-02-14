package controllers

import (
	"formation-go/lib"
	"formation-go/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func CreatePersonController(context echo.Context) error {
	personContext, err := lib.BindAndValidate[requests.PersonRequest](context)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = lib.CreatePerson(personContext)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, gin.H{"OK": true})
}
