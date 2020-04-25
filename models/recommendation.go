package models

import (
	"database/sql"
	"math/rand"
)

func getRandom(ec EntertainmentCollection) Entertainment {
	ri := rand.Intn(len(ec.Entertainments))

	e := ec.Entertainments[ri]

	return e
}

func GetRecommendation(db *sql.DB) Entertainment {
	sql := "SELECT e.id, e.name, e.rating FROM entertainment e LEFT JOIN watched w ON(e.id = w.entertainment_id) WHERE w.id IS NULL"
	return getRandom(GetEntertainments(db, sql))
}

func GetMovieRecommendation(db *sql.DB) Entertainment {
	return getRandom(GetUnwatchedMovies(db))
}

func GetSerieRecommendation(db *sql.DB) Entertainment {
	return getRandom(GetUnwatchedSeries(db))
}
