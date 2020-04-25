package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ezeoleaf/watch-this/models"
	"github.com/labstack/echo"
)

func GetSeries(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetSeries(db))
	}
}
