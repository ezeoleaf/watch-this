package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/ezeoleaf/watch-this/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbName := os.Getenv("DATABASE_NAME")

	db := initDB(dbName)
	migrate(db)

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	initRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}

func initRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to WatchThis")
	})
	e.GET("/recommendation", handlers.GetRecommendation(db))
	e.GET("/movies", handlers.GetMovies(db))
	e.GET("/series", handlers.GetSeries(db))
	e.GET("/recommendation/movie", handlers.GetMovieRecommendation(db))
	e.GET("/recommendation/serie", handlers.GetSerieRecommendation(db))
	e.GET("/rating/:name", handlers.GetRating(db))
	e.GET("/types", handlers.GetTypes(db))

	e.POST("/entertainment", handlers.PostEntertainment(db))
	e.POST("/movie", handlers.PostMovie(db))
	e.POST("/serie", handlers.PostSerie(db))
	e.POST("/user", handlers.PostUser(db))
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
	sql := `
	CREATE TABLE IF NOT EXISTS type(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS entertainment(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		rating DOUBLE,
		type_id INTEGER REFERENCES type(id)
	);
	CREATE TABLE IF NOT EXISTS user(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username VARCHAR NOT NULL,
		email VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS wish_list(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL REFERENCES user(id),
		entertainment_id INTEGER NOT NULL REFERENCES entertainment(id),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS recommendation(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL REFERENCES user(id),
		entertainment_id INTEGER NOT NULL REFERENCES entertainment(id),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS watched(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		entertainment_id INTEGER NOT NULL REFERENCES entertainment(id),
		user_id INTEGER NOT NULL REFERENCES user(id),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	--INSERT INTO type(id, name) VALUES(1, 'Movie'), (2, 'Serie');
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
