package tasksvc

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/columnsvc"
	"github.com/go-playground/validator/v10"
	"sync"
)

type service struct {
	storage  TaskStorage
	validate *validator.Validate
}

var once sync.Once
var singleton service

func NewService(cs TaskStorage) {
	once.Do(func() {
		singleton.storage = cs
		singleton.validate = validator.New()
	})
}

func validate(t *models.Task) error {
	return singleton.validate.Struct(t)
}

func Create(t *models.Task) error {
	err := validate(t)
	if err != nil {
		return err
	}

	return singleton.storage.Create(t)
}

func Delete(taskId int) (*models.Task, error) {
	t, err := singleton.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	err = singleton.storage.Delete(t)
	return t, err
}

func GetById(taskId int) (*models.Task, error) {
	return singleton.storage.GetById(taskId)
}

func GetListByColumnId(taskId int) (*[]models.Task, error) {
	return singleton.storage.GetListByColumnId(taskId)
}

func Move(taskId int, cp *models.TaskPosition) (*models.Task, error) {
	c, err := singleton.storage.GetById(taskId)
	if err != nil {
		return c, err
	}

	if cp.Position == c.Position {
		return c, nil
	}

	err = singleton.storage.Move(c, cp.Position)
	return c, err
}

func Shift(taskId int, tc *models.TaskColumn) (*models.Task, error) {
	t, err := singleton.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	if tc.ColumnId == t.ColumnId {
		return t, nil
	}

	inSameProject, err := columnsvc.IsSameProject(t.ColumnId, tc.ColumnId)
	if err != nil {
		return t, err
	}
	if !inSameProject {
		return t, fmt.Errorf("columns must be in the same project")
	}

	err = singleton.storage.Shift(t, tc.ColumnId)
	return t, err
}

func Update(taskId int, tb *models.TaskBase) (*models.Task, error) {
	t, err := singleton.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	t.TaskBase = *tb
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
