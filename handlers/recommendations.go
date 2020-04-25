package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ezeoleaf/watch-this/models"

	"github.com/labstack/echo"
)

func GetRecommendation(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetRecommendation(db))
	}
}

func GetMovieRecommendation(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetMovieRecommendation(db))
	}
}

func GetSerieRecommendation(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetSerieRecommendation(db))
	}
}
