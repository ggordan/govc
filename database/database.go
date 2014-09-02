package database

import (
	"database/sql"

	"github.com/ggordan/govc/conf"
	_ "github.com/mattn/go-sqlite3"
)

// GetDB returns the database, and creates it if it doesn't exist yet.
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", conf.GetDBPath())
	if err != nil {
		return nil, err
	}
	return db, nil
}
