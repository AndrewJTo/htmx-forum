package migrations

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func as() {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, email VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, isAdmin BOOLEAN NOT NULL CHECK (mycolumn IN (0, 1) DEFAULT 0), isBanned BOOLEAN NOT NULL CHECK (mycolumn IN (0, 1) DEFAULT 0))")
	if err != nil {
		panic(err)
	}
}
