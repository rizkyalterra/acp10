package middlewares

import (
	"acp10/configs"
	"acp10/models/users"

	"github.com/labstack/echo/v4"
)

func BasicAuthMiddlewares(email, password string, c echo.Context) (bool, error) {
	var user users.User
	err := configs.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
