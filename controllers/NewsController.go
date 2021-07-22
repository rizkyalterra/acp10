package controllers

import (
	"acp10/configs"
	"acp10/lib/database"
	"acp10/models/news"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewsControllers(c echo.Context) error {

	var newsCreate news.NewsCreate
	c.Bind(&newsCreate)

	var newsDB news.News
	newsDB.Title = newsCreate.Title
	newsDB.Content = newsCreate.Content

	err := configs.DB.Create(&newsDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newsDB)
}

func GetNewsControllers(c echo.Context) error {

	var newsData []news.News
	var err error
	newsData, err = database.GetDataUserAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			newsData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data News",
		newsData,
	))
}
