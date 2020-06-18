package columnsvc

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type service struct {
	storage ColumnStorage
	validate *validator.Validate
}

var once sync.Once
var singleton service

func NewService(cs ColumnStorage) {
	once.Do(func() {
		singleton.storage = cs
		singleton.validate = validator.New()
	})
}

func validate(column *models.Column) error {
	return singleton.validate.Struct(column)
}

func Create(c *models.Column) error {
	err := validate(c)
	if err != nil {
		return err
	}

	return singleton.storage.Create(c)
}

func Delete(columnId int) (*models.Column, error) {
	c, err := singleton.storage.GetById(columnId)
	if err != nil {
		return c, err
	}

	count, err := singleton.storage.GetCountsByProjectId(c.ProjectId)
	if err != nil {
		return c, err
	}

	if count == 1 {
		return c, fmt.Errorf("deleting last column is not allowed")
	}
	err = singleton.storage.Delete(c)
	return c, err
}

func GetById(columnId int) (*models.Column, error) {
	return singleton.storage.GetById(columnId)
}

func GetListByProjectId(projectId int) (*[]models.Column, error) {
	return singleton.storage.GetListByProjectId(projectId)
}

func Move(columnId int, cp *models.ColumnPosition) (*models.Column, error){
	c, err := singleton.storage.GetById(columnId)
	if err != nil {
		return c, err
	}

	if cp.Position == c.Position {
		return c, nil
	}
	c.ColumnPosition = *cp

	err = singleton.storage.Move(c)
	return c, err
}

func Update(columnId int, cb *models.ColumnBase) (*models.Column, error) {
	c, err := singleton.storage.GetById(columnId)
	if err != nil {
		return c, err
	}

	c.ColumnBase = *cb
	err = validate(c)
	if err != nil {
		return c, err
	}

	err = singleton.storage.Update(c)
	return c, err
}
