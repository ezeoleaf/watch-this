package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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

func PostWatched(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		entertainmentId, _ := strconv.Atoi(c.FormValue("entertainmentId"))
		userId, _ := strconv.Atoi(c.FormValue("userId"))

		id, err := models.PostWatched(db, userId, entertainmentId)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"created": id,
		})
	}
}
