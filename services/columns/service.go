package columns

import (
	"sync"
)

type service struct {
	storage ColumnStorage
}

var once sync.Once
var singleton service

func NewService(cs ColumnStorage) {
	once.Do(func() {
		singleton.storage = cs
	})
}

func Storage() ColumnStorage {
	return singleton.storage
}
