package models

import "database/sql"

type EntertainmentType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EntertainmentTypeCollection struct {
	Types []EntertainmentType `json:"types"`
}

func GetTypes(db *sql.DB) EntertainmentTypeCollection {
	sql := "SELECT id, name FROM type"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := EntertainmentTypeCollection{}

	for rows.Next() {
		t := EntertainmentType{}
		e := rows.Scan(&t.ID, &t.Name)
		if e != nil {
			panic(e)
		}
		result.Types = append(result.Types, t)
	}

	return result
}
