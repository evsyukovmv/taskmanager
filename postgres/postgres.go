package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

var once sync.Once
var db *sql.DB

func Configure(databaseURL string) error {
	var err error
	once.Do(func() {
		db, err = sql.Open("postgres", databaseURL)
	})

	if err != nil {
		return err
	}

	return db.Ping()
}

func DB() *sql.DB {
	return db
}
