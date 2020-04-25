package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ezeoleaf/watch-this/models"
	"github.com/labstack/echo"
)

// H is a type used to return data
type H map[string]interface{}

// PostUser saves an user to the database
func PostUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User

		c.Bind(&user)

		id, err := models.PostUser(db, user.Username, user.Email)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{
			"created": id,
		})
	}
}
