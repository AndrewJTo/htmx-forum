package migrations

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *dbx.DB) {
	driver, err := sqlite.WithInstance(db.DB(), &sqlite.Config{})
	if err != nil {
		panic("Could not create DB driver for migrations: " + err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite", driver)
	if err != nil {
		panic("Could create migration instance: " + err.Error())
	}
	err = m.Up()
	if err != nil {
		panic("Could run migrations: " + err.Error())
	}
}
