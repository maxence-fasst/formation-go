package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheckRoute(context echo.Context) error {
	return context.JSON(http.StatusOK, gin.H{"ok": true})
}
