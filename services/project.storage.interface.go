package services

import "github.com/evsyukovmv/taskmanager/models"

type ProjectStorage interface {
	GetList() (projects *[]models.Project, err error)
	GetById(id int) (*models.Project, error)
	Create(project *models.Project) error
	Update(project *models.Project) error
	Delete(project *models.Project) error
	Clear() error
}
