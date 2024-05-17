package core

import (
	"database/sql"
)

// Database is an interface for database connection.
type Database interface {
	Connect() (*sql.DB, error)
}
