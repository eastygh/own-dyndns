package db

import (
	"database/sql"
	"embed"
	_ "github.com/glebarez/go-sqlite"
	"github.com/pressly/goose/v3"
	_ "github.com/pressly/goose/v3"
	"src/config"
)

var dbInstance *sql.DB

//go:embed migrations-sqlite/*.sql
var embedMigrations embed.FS

func Init(cfg *config.DB) *sql.DB {
	db, err := sql.Open("sqlite", cfg.Url)
	if err != nil {
		panic(err)
	}
	set(db)
	return Get()
}

func ApplyMigrations(cfg *config.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(cfg.Type); err != nil {
		panic(err)
	}

	if err := goose.Up(Get(), "migrations-sqlite"); err != nil {
		panic(err)
	}
}

func set(db *sql.DB) {
	dbInstance = db
}
func Get() *sql.DB {
	return dbInstance
}
