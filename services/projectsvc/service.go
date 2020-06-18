package projectsvc

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type service struct {
	storage ProjectStorage
	validate *validator.Validate
}

var once sync.Once
var singleton service

func NewService(ps ProjectStorage) {
	once.Do(func() {
		singleton.storage = ps
		singleton.validate = validator.New()
	})
}

func validate(column *models.Project) error {
	return singleton.validate.Struct(column)
}

func Create(p *models.Project) error {
	err := validate(p)
	if err != nil {
		return err
	}

	return singleton.storage.Create(p)
}

func Delete(projectId int) (*models.Project, error) {
	p, err := singleton.storage.GetById(projectId)
	if err != nil {
		return p, err
	}

	err = singleton.storage.Delete(p)
	return p, err
}

func GetById(projectId int) (*models.Project, error) {
	return singleton.storage.GetById(projectId)
}

func GetList() (*[]models.Project, error) {
	return singleton.storage.GetList()
}

func Update(projectId int, pb *models.ProjectBase) (*models.Project, error) {
	p, err := singleton.storage.GetById(projectId)
	if err != nil {
		return p, err
	}

	p.ProjectBase = *pb
	err = validate(p)
	if err != nil {
		return p, err
	}

	err = singleton.storage.Update(p)
	return p, err
}
