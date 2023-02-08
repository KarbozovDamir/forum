package sqlworker

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type SQL struct {
	Server *sql.DB
	Repo   string
}
