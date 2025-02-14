package apiHandler

import (
	"formation-go/controllers"

	"github.com/labstack/echo/v4"
)

// @Param request body requests.PersonRequest true "Person body"
// @Success 200
// @Router /person [post]
func CreatePerson(context echo.Context) error {
	return controllers.CreatePersonController(context)
}
