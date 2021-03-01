package db

import (
	"database/sql"
)

type Store interface {
	open() error
	QueryRow(query string, args ...interface{}) *sql.Row
}
