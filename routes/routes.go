package routes

import (
	"acp10/constants"
	"acp10/controllers"
	"acp10/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddlewares(e)

	eAuth := e.Group("")
	// eBasicAuth.Use(middleware.BasicAuth(middlewares.BasicAuthMiddlewares))
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))
	eAuth.GET("/news", controllers.GetNewsControllers)
	eAuth.POST("/news", controllers.CreateNewsControllers)
	eAuth.GET("/news/:newsId", controllers.DetailNewsControllers)

	e.POST("/login", controllers.LoginController)

	return e
}
