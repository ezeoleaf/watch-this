package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ezeoleaf/watch-this/models"
	"github.com/labstack/echo"
)

func GetRating(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")

		return c.JSON(http.StatusOK, models.GetRating(db, name))
	}
}

func PostEntertainment(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		userID, _ := strconv.Atoi(c.FormValue("user_id"))
		titleTypeID, _ := strconv.Atoi(c.FormValue("type_id"))

		id, err := models.PostEntertainment(db, title, userID, titleTypeID)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"created": id,
		})
	}
}
