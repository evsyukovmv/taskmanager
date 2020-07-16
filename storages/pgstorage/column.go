package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"strconv"
)

type PostgresColumnsStorage struct{}

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
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.ProjectId); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO columns (name, project_id) VALUES ($1, $2) RETURNING id`)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(column.Name, column.ProjectId).Scan(&column.Id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (c *PostgresColumnsStorage) Move(column *models.Column, newPosition int) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	{
		querySetToNull := `UPDATE columns SET position = NULL WHERE id = $1`
		stmt, err := tx.Prepare(querySetToNull)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.Id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryDecrement := `UPDATE columns SET position = position - 1 WHERE project_id = $1 AND position > $2`
		stmt, err := tx.Prepare(queryDecrement)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.ProjectId, column.Position); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryIncrement := `UPDATE columns SET position = position + 1 WHERE project_id = $1 AND position >= $2`
		stmt, err := tx.Prepare(queryIncrement)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.ProjectId, newPosition); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryUpdatePosition := `UPDATE columns SET position = $2 WHERE id = $1`
		stmt, err := tx.Prepare(queryUpdatePosition)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.Id, newPosition); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	column.Position = newPosition

	return nil
}

func (c *PostgresColumnsStorage) Update(column *models.Column) error {
	sqlUpdate := `UPDATE columns SET name = $2 WHERE id = $1`
	_, err := postgres.DB().Exec(sqlUpdate, column.Id, column.Name)
	return err
}

func (c *PostgresColumnsStorage) Delete(column *models.Column) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	var nearColumnId int
	var maxNearPosition int

	{
		var leftOrRight string
		if column.Position == 0 {
			leftOrRight = `>`
		} else {
			leftOrRight = `<`
		}

		queryNearColumn := `SELECT id FROM columns WHERE project_id = $1 AND position ` + leftOrRight + ` $2 ORDER BY position ASC LIMIT 1`

		stmt, err := tx.Prepare(queryNearColumn)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(column.ProjectId, column.Position).Scan(&nearColumnId); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryMaxLeftPosition := `SELECT coalesce(max(position), -1) FROM tasks WHERE column_id = $1`
		stmt, err := tx.Prepare(queryMaxLeftPosition)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(nearColumnId).Scan(&maxNearPosition); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		sqlMove := `UPDATE tasks SET column_id = $1, position = position + $3 WHERE column_id = $2`
		stmt, err := tx.Prepare(sqlMove)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(nearColumnId, column.Id, maxNearPosition+1); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		sqlDelete := `DELETE FROM columns WHERE id = $1`
		stmt, err := tx.Prepare(sqlDelete)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(column.Id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (c *PostgresColumnsStorage) InSameProject(columnIds ...int) (bool, error) {
	var count int
	ids := ""
	for _, v := range columnIds {
		if len(ids) > 0 {
			ids += ","
		}
		ids += strconv.Itoa(v)
	}
	query := `SELECT COUNT(DISTINCT(project_id)) FROM columns WHERE id IN (` + ids + `)`
	if err := postgres.DB().QueryRow(query).Scan(&count); err != nil {
		return false, err
	}

	return count == 1, nil
}

func (c *PostgresColumnsStorage) CountInProject(projectId int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM columns WHERE project_id = $1`
	err := postgres.DB().QueryRow(query, projectId).Scan(&count)
	return count, err
}

func (c *PostgresColumnsStorage) Clear() error {
	_, err := postgres.DB().Exec("TRUNCATE columns RESTART IDENTITY CASCADE;")
	return err
}
