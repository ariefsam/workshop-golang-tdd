package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InsertPerson(firstName, lastName string) (err error) {
	db, err := sql.Open("sqlite3", "name.db")
	if err != nil {
		return
	}
	db.SetMaxOpenConns(1)

	sqlStmt := `
	INSERT INTO person (first_name, last_name) VALUES (?,?);
	
	`
	_, err = db.Exec(sqlStmt, firstName, lastName)
	return
}
