package daos

import "database/sql"

type Dao struct {
	DB *sql.DB
}
