package postgres

import (
	"github.com/go-pg/pg/v9"
	"sync"
)

var once sync.Once
var db *pg.DB

func Configure(host, port, user, password, database string) error {
	var err error
	once.Do(func() {
		db = pg.Connect(&pg.Options{
			Addr:     host + ":" + port,
			User:     user,
			Password: password,
			Database: database,
		})
	})

	_, err = db.Exec("SELECT 1")
	return err
}

func DB() *pg.DB {
	return db
}
