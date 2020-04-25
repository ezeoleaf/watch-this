package models

import "database/sql"

const Movie = "Movie"

func GetMovies(db *sql.DB) EntertainmentCollection {
	sql := "SELECT e.id, e.name, e.rating FROM entertainment e INNER JOIN type t ON(t.id = e.type_id AND t.name = '" + Movie + "')"
	return GetEntertainments(db, sql)
}

func GetUnwatchedMovies(db *sql.DB) EntertainmentCollection {
	sql := "SELECT e.id, e.name, e.rating FROM entertainment e INNER JOIN type t ON(t.id = e.type_id AND t.name = '" + Movie + "') LEFT JOIN watched w ON(e.id = w.entertainment_id) WHERE w.id IS NULL"
	return GetEntertainments(db, sql)
}
