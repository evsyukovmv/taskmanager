package services

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type ProjectService struct {
	storage   ProjectStorage
	validator *validator.Validate
}

var projectOnce sync.Once
var projectSingleton *ProjectService

func NewProjectService(ps ProjectStorage) {
	projectOnce.Do(func() {
		projectSingleton = &ProjectService{storage: ps, validator: validator.New()}
	})
}

func ForProject() *ProjectService {
	return projectSingleton
}

func (s *ProjectService) validate(project *models.Project) error {
	return projectSingleton.validator.Struct(project)
}

func (s *ProjectService) Create(p *models.Project) error {
	err := s.validate(p)
	if err != nil {
		return err
	}

	return s.storage.Create(p)
}

func (s *ProjectService) Delete(projectId int) (*models.Project, error) {
	p, err := s.GetById(projectId)
	if err != nil {
		return p, err
	}

	err = s.storage.Delete(p)
	return p, err
}

func (s *ProjectService) GetById(projectId int) (*models.Project, error) {
	return s.storage.GetById(projectId)
}

func (s *ProjectService) GetList() (*[]models.Project, error) {
	return s.storage.GetList()
}

func (s *ProjectService) Update(projectId int, pb *models.ProjectBase) (*models.Project, error) {
	p, err := s.storage.GetById(projectId)
	if err != nil {
		return p, err
	}

	p.ProjectBase = *pb
	err = s.validate(p)
	if err != nil {
		return p, err
	}

	err = s.storage.Update(p)
	return p, err
}

func (s *ProjectService) Clear() error {
	return s.storage.Clear()
}
