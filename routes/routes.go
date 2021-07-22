package routes

import (
	"acp10/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/news", controllers.GetNewsControllers)
	e.POST("/news", controllers.CreateNewsControllers)
	return e
}
