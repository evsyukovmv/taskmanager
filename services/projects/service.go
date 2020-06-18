package projects

import (
	"sync"
)

type service struct {
	storage ProjectStorage
}

var once sync.Once
var singleton service

func NewService(ps ProjectStorage) {
	once.Do(func() {
		singleton.storage = ps
	})
}

func Storage() ProjectStorage {
	return singleton.storage
}
