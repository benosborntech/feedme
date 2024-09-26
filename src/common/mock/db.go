package mock

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewMockDB(schema string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
