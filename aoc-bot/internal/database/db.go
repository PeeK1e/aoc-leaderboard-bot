package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var dbInstance *sql.DB

func Open(path string) error {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	dbInstance = db
	return initDb()
}

func initDb() error {
	q := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, day INTEGER, stars INTEGER);`
	_, err := dbInstance.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = dbInstance.Close()
}

func RunQuery(s string, args ...any) (sql.Result, error) {
	tx, err := dbInstance.Begin()
	if err != nil {
		return nil, err
	}

	r, err := dbInstance.Exec(s, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return r, nil
}
