package models

import "database/sql"

const Serie = "Serie"

func GetSeries(db *sql.DB) EntertainmentCollection {
	sql := "SELECT e.id, e.name, e.rating FROM entertainment e INNER JOIN type t ON(t.id = e.type_id AND t.name = '" + Serie + "')"
	return GetEntertainments(db, sql)
}

func GetUnwatchedSeries(db *sql.DB) EntertainmentCollection {
	sql := "SELECT e.id, e.name, e.rating FROM entertainment e INNER JOIN type t ON(t.id = e.type_id AND t.name = '" + Serie + "') LEFT JOIN watched w ON(e.id = w.entertainment_id) WHERE w.id IS NULL"
	return GetEntertainments(db, sql)
}
