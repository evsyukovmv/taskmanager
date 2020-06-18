package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
)

type PostgresProjectsStorage struct {}

func (p *PostgresProjectsStorage) GetList() (*[]models.Project, error) {
	var projects []models.Project
	err := postgres.DB().Model(&projects).Order("name ASC").Select()
	return &projects, err
}

func (p *PostgresProjectsStorage) GetByID(id int) (*models.Project, error) {
	project := &models.Project{Id: id}
	err := postgres.DB().Select(project)
	return project, err
}

func (p *PostgresProjectsStorage) Create(project *models.Project) error {
	err := postgres.DB().Insert(project)
	if err != nil {
		return err
	}
	column := &models.Column{ColumnBase: models.ColumnBase{ Name: "Default" }, ProjectId: project.Id}
	err = postgres.DB().Insert(column)
	return err
}

func (p *PostgresProjectsStorage) Update(project *models.Project) error {
	return postgres.DB().Update(project)
}

func (p *PostgresProjectsStorage) Delete(project *models.Project) error {
	return postgres.DB().Delete(project)
}
