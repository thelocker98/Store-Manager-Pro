package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(1) // SQLite prefers 1 writer
	DB.SetMaxIdleConns(1)

	createTables()
}
