package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
)

type PostgresColumnsStorage struct {}

func (c *PostgresColumnsStorage) GetCountsByProjectId(projectId int) (count int, err error) {
	row := postgres.DB().QueryRow("SELECT COUNT (*) FROM columns WHERE project_id = $1", projectId)
	err = row.Scan(&count)
	return
}

func (c *PostgresColumnsStorage) GetListByProjectId(projectId int) (*[]models.Column, error) {
	var columns []models.Column

	rows, err := postgres.DB().Query("SELECT * FROM columns WHERE project_id = $1 ORDER BY position ASC", projectId)
	if err != nil {
		return &columns, err
	}

	for rows.Next() {
		var column models.Column

		if err := rows.Scan(&column.Id, &column.Name, &column.Position, &column.ProjectId); err != nil {
			return &columns, err
		}

		columns = append(columns, column)
	}

	return &columns, err
}

func (c *PostgresColumnsStorage) GetById(id int) (*models.Column, error) {
	column := &models.Column{Id: id}
	err := postgres.DB().QueryRow(
		"SELECT * FROM columns WHERE id = $1", column.Id).Scan(
			&column.Id, &column.Name, &column.Position, &column.ProjectId)
	return column, err
}

func (c *PostgresColumnsStorage) Create(column *models.Column) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	{
		stmt, err := tx.Prepare(`UPDATE columns SET position = position + 1 WHERE project_id = $1`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.ProjectId); err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO columns (name, project_id) VALUES ($1, $2) RETURNING id`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(column.Name, column.ProjectId).Scan(&column.Id); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (c *PostgresColumnsStorage) Move(column *models.Column) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	{
		stmt, err := tx.Prepare(`UPDATE columns SET position = NULL WHERE id = $1`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.Id); err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare(`UPDATE columns SET position = position + 1 WHERE project_id = $1`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.ProjectId); err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare(`UPDATE columns SET position = $1 WHERE id = $2`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.Position, column.Id); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (c *PostgresColumnsStorage) Update(column *models.Column) error {
	sqlUpdate := `UPDATE columns SET name = $2 WHERE id = $1`
	_, err := postgres.DB().Exec(sqlUpdate, column.Id, column.Name)
	return err
}

func (c *PostgresColumnsStorage) Delete(column *models.Column) error {
	sqlDelete := `DELETE FROM columns WHERE id = $1;`
	_, err := postgres.DB().Exec(sqlDelete, column.Id)
	return err
}
