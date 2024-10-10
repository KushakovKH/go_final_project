package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitiDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL CHECK(LENGTH(date) = 8),
		title TEXT NOT NULL CHECK(LENGTH(title) <= 255),
		comment TEXT TEXT CHECK(LENGTH(comment) <= 1000),
		repeat TEXT CHECK(LENGTH(repeat) <= 20)
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
