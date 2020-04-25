package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ezeoleaf/watch-this/models"
	"github.com/labstack/echo"
)

func GetRating(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")

		return c.JSON(http.StatusOK, models.GetRating(db, name))
	}
}
