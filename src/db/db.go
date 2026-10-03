package db

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

//go:embed migrations/*.sql
var migrations embed.FS

func InitDB(path string) {
	var err error

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "sqlite", "":
		databaseFilePath := filepath.Join(path, "store.db")

		DB, err = sql.Open("sqlite3", databaseFilePath)
		if err != nil {
			panic(err)
		}

		DB.SetMaxOpenConns(1)
		DB.SetMaxIdleConns(1)

		if err := goose.SetDialect("sqlite"); err != nil {
			panic(err)
		}

	default:
		panic("unsupported DB_TYPE: " + dbType)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	// Run Migrations
	goose.SetBaseFS(migrations)

	if err := goose.Up(DB, "migrations"); err != nil {
		panic(err)
	}
}
