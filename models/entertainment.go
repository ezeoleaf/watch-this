package models

import (
	"database/sql"
	"strings"
)

type Entertainment struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

type EntertainmentCollection struct {
	Entertainments []Entertainment `json:"items"`
}

func GetAllEntertainments(db *sql.DB) EntertainmentCollection {
	query := "SELECT e.id, e.name, e.rating FROM entertainment e"
	return GetEntertainments(db, query)
}

func GetEntertainments(db *sql.DB, sql string) EntertainmentCollection {
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := EntertainmentCollection{}
	for rows.Next() {
		entertainment := Entertainment{}
		e := rows.Scan(&entertainment.ID, &entertainment.Name, &entertainment.Rating)
		if e != nil {
			panic(e)
		}

		result.Entertainments = append(result.Entertainments, entertainment)
	}

	return result
}

func GetEntertainmentByName(db *sql.DB, title string) int64 {
	query := "SELECT e.id FROM entertainment e WHERE upper(e.name) = upper($1)"

	var id int64

	row := db.QueryRow(query, title)
	err := row.Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return id
}

func PostWatched(db *sql.DB, userID int, entertainmentID int) (int64, error) {
	query := "INSERT INTO watched(user_id, entertainment_id, created_at) VALUES(?, ?, now())"

	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	r, e := stmt.Exec(userID, entertainmentID)

	if e != nil {
		panic(e)
	}

	return r.LastInsertId()
}

func PostEntertainment(db *sql.DB, title string, userID int, titleTypeId int) (int64, error) {
	id := GetEntertainmentByName(db, title)

	if id != 0 {
		return id, nil
	}

	// Need to create a new entertainment
	title = strings.Title(strings.ToLower(title))
	rating := GetRating(db, title)

	query := "INSERT INTO entertainment(name, rating, type_id) VALUES(?, ?, ?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	r, e := stmt.Exec(title, rating, titleTypeId)

	if e != nil {
		panic(e)
	}

	id, err = r.LastInsertId()

	if err != nil {
		panic(err)
	}

	//Associate user to entertainment

	return id, err
}
