package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// OMDB is a struct with a response from OMDB API
type OMDB struct {
	Title  string `json:"Title"`
	Rating string `json:"imdbRating"`
}

// GetRating returns the IMDB rating for a movie or serie title
func GetRating(db *sql.DB, title string) float64 {

	u := os.Getenv("OMDB_API_URL")
	key := os.Getenv("OMDB_API_KEY")

	u = strings.ReplaceAll(u, "{apikey}", key)
	u = strings.ReplaceAll(u, "{title}", url.QueryEscape(title))

	response, err := http.Get(u)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	r := &OMDB{}
	err = json.Unmarshal(body, r)

	if err != nil {
		return 0
	}

	rating, e := strconv.ParseFloat(r.Rating, 64)

	if e != nil {
		return 0
	}

	return rating
}
