package models

import "database/sql"

// User represents a user stored in the database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// PostUser creates a new record in user table and returns the id of the created user
func PostUser(db *sql.DB, username string, email string) (int64, error) {
	sql := "INSERT INTO user(username, email) VALUES(?, ?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	r, err := stmt.Exec(username, email)

	if err != nil {
		panic(err)
	}

	return r.LastInsertId()
}
