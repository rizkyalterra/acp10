package controllers

import (
	"acp10/lib/database"
	"acp10/middlewares"
	"acp10/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {

	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	userDB, e := database.LoginUser(userLogin)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}

	token, _ := middlewares.GenerateTokenJWT(userDB.Id)

	var userResponse = users.UserResponse{
		Id:        userDB.Id,
		Name:      userDB.Name,
		Email:     userDB.Name,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: userDB.UpdatedAt,
		DeletedAt: userDB.DeletedAt,
		Token:     token,
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		userResponse,
	))

}
