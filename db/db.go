package db

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(path string) {
	// Create file
	databaseFilePath := filepath.Join(path, "store.db")

	var err error
	DB, err = sql.Open("sqlite3", databaseFilePath)
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(1) // SQLite prefers 1 writer
	DB.SetMaxIdleConns(1)

	createTables()
}
