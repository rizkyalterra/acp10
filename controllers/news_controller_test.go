package controllers

import (
	"acp10/configs"
	"acp10/models/news"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	if Addition(2, 3) != 5 {
		t.Error("Salah harusnya 2 + 3 adalah 5")
	}
	if Addition(-1, -2) != 0 {
		t.Error("Salah harusnya -1 + -2 adalah 0")
	}
}

// setup echo
func InitEcho() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func CreateSeedNews() {
	var newsDB news.News
	newsDB.Title = "Test news"
	newsDB.Content = "Test enews Content"
	configs.DB.Create(&newsDB)
}

func TestGetNewsControllers(t *testing.T) {
	e := InitEcho()
	CreateSeedNews()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/news")

	if assert.NoError(t, GetNewsControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := news.NewsResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
		assert.Equal(t, 2, len(baseResponse.Data))
	}
}
