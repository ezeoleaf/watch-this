package models

import (
	"database/sql"
	"fmt"
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
	sql := "SELECT e.id, e.name, e.rating FROM entertainment e"
	return GetEntertainments(db, sql)
}

func GetEntertainments(db *sql.DB, sql string) EntertainmentCollection {
	fmt.Println(sql)
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
