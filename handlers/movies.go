package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ezeoleaf/watch-this/models"
	"github.com/labstack/echo"
)

func GetMovies(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetMovies(db))
	}
}

func PostMovie(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		userID, _ := strconv.Atoi(c.FormValue("user_id"))

		id, err := models.PostEntertainment(db, title, userID, 1)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"created": id,
		})
	}
}
