package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableName() (err error) {
	db, err := sql.Open("sqlite3", "name.db")
	if err != nil {
		return
	}
	db.SetMaxOpenConns(1)

	sqlStmt := `
	CREATE TABLE if not exists person (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL DEFAULT "",
		last_name TEXT NOT NULL DEFAULT "",
		created_at NUMERIC NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at NUMERIC NOT NULL DEFAULT 0);
	
	`
	_, err = db.Exec(sqlStmt)
	return
}
