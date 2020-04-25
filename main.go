package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("watch.db")
	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	initRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}

func initRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to WatchThis")
	})
	e.GET("/recomendation", handlers.GetRecomendation(db))
	e.GET("/movies", handlers.GetMovies(db))
	e.GET("/series", handlers.GetSeries(db))
	e.GET("/recomendation/movie", handlers.GetMovieRecomendation(db))
	e.GET("/recomendation/serie", handlers.GetSerieRecomendation(db))
	e.GET("/rating/:name", handlers.GetRating(db))

	e.POST("/recomendation", handlers.PostRecomendation(db))
	e.POST("/watched", handlers.PostWatched(db))

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("There is no db")
	}

	return db
}

func migrate(db *sql.DB) {

}
