package db

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

//go:embed migrations/*.sql
var migrations embed.FS

func InitDB(path string) {
	var err error

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres", "":
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		connString := `postgres://` + user + `:` + pass + `@` + host + `:` + port + `/` + dbname + `?sslmode=disable`

		DB, err = sql.Open("postgres", connString)
		if err != nil {
			panic(err)
		}

		DB.SetMaxOpenConns(10)
		DB.SetMaxIdleConns(5)

	case "sqlite":
		databaseFilePath := filepath.Join(path, "store.db")

		DB, err = sql.Open("sqlite3", databaseFilePath)
		if err != nil {
			panic(err)
		}

		DB.SetMaxOpenConns(1)
		DB.SetMaxIdleConns(1)

	default:
		panic("unsupported DB_TYPE: " + dbType)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	// Run Migrations
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(DB, "migrations"); err != nil {
		panic(err)
	}
}
