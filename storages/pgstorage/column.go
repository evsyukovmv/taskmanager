package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-pg/pg/v9"
)

type PostgresColumnsStorage struct {}

func (c *PostgresColumnsStorage) GetCountsByProjectId(projectId int) (count int, err error) {
	_, err = postgres.DB().QueryOne(pg.Scan(&count), "SELECT COUNT (*) FROM columns WHERE project_id = ?", projectId)
	return
}

func (c *PostgresColumnsStorage) GetListByProjectId(projectId int) (*[]models.Column, error) {
	var columns []models.Column

	err := postgres.DB().Model(&columns).Where(
		"project_id = ?", projectId).Order("position ASC").Select()
	return &columns, err
}

func (c *PostgresColumnsStorage) GetById(id int) (*models.Column, error) {
	column := &models.Column{Id: id}
	err := postgres.DB().Select(column)
	return column, err
}

func (c *PostgresColumnsStorage) Create(column *models.Column) error {
	return postgres.DB().Insert(column)
}

func (c *PostgresColumnsStorage) Move(column *models.Column) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	_, err = postgres.DB().Exec("UPDATE columns SET position = NULL WHERE id = ?", column.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = postgres.DB().Exec(
		"UPDATE columns SET position = position + 1 WHERE project_id = ?",
		column.ProjectId, column.Position, column.Id,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = postgres.DB().Exec("UPDATE columns SET position = ? WHERE id = ?", column.Position, column.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (c *PostgresColumnsStorage) Update(column *models.Column) error {
	return postgres.DB().Update(column)
}

func (c *PostgresColumnsStorage) Delete(project *models.Column) error {
	return postgres.DB().Delete(project)
}
