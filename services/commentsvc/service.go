package commentsvc

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type service struct {
	storage CommentStorage
	validate *validator.Validate
}

var once sync.Once
var singleton service

func NewService(cs CommentStorage) {
	once.Do(func() {
		singleton.storage = cs
		singleton.validate = validator.New()
	})
}

func validate(t *models.Comment) error {
	return singleton.validate.Struct(t)
}

func Create(t *models.Comment) error {
	err := validate(t)
	if err != nil {
		return err
	}

	return singleton.storage.Create(t)
}

func Delete(taskId int) (*models.Comment, error) {
	t, err := singleton.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	err = singleton.storage.Delete(t)
	return t, err
}

func GetById(taskId int) (*models.Comment, error) {
	return singleton.storage.GetById(taskId)
}

func GetListByTaskId(taskId int) (*[]models.Comment, error) {
	return singleton.storage.GetListByTaskId(taskId)
}

func Update(taskId int, tb *models.CommentBase) (*models.Comment, error) {
	t, err := singleton.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	t.CommentBase = *tb
	err = validate(t)
	if err != nil {
		return t, err
	}

	err = singleton.storage.Update(t)
	return t, err
}

func Clear() error {
	return singleton.storage.Clear()
}
