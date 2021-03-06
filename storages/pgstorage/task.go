package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
)

type PostgresTasksStorage struct{}

func (c *PostgresTasksStorage) GetListByColumnId(columnId int) (*[]models.Task, error) {
	var tasks []models.Task

	rows, err := postgres.DB().Query("SELECT * FROM tasks WHERE column_id = $1 ORDER BY position", columnId)
	if err != nil {
		return &tasks, err
	}

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Position, &task.ColumnId); err != nil {
			return &tasks, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, err
}

func (c *PostgresTasksStorage) GetById(id int) (*models.Task, error) {
	task := &models.Task{Id: id}
	err := postgres.DB().QueryRow(
		"SELECT * FROM tasks WHERE id = $1", task.Id).Scan(
		&task.Id, &task.Name, &task.Description, &task.Position, &task.ColumnId)
	return task, err
}

func (c *PostgresTasksStorage) Create(task *models.Task) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	{
		stmt, err := tx.Prepare(`UPDATE tasks SET position = position + 1 WHERE column_id = $1`)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(task.ColumnId); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO tasks (name, description, column_id) VALUES ($1, $2, $3) RETURNING id`)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(task.Name, task.Description, task.ColumnId).Scan(&task.Id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (c *PostgresTasksStorage) Move(task *models.Task, newPosition int) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	{
		querySetToNull := `UPDATE tasks SET position = NULL WHERE id = $1`
		stmt, err := tx.Prepare(querySetToNull)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(task.Id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryDecrement := `UPDATE tasks SET position = position - 1 WHERE column_id = $1 AND position > $2`
		stmt, err := tx.Prepare(queryDecrement)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(task.ColumnId, task.Position); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryIncrement := `UPDATE tasks SET position = position + 1 WHERE column_id = $1 AND position >= $2`
		stmt, err := tx.Prepare(queryIncrement)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(task.ColumnId, newPosition); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		queryUpdatePosition := `UPDATE tasks SET position = $2 WHERE id = $1`
		stmt, err := tx.Prepare(queryUpdatePosition)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(task.Id, newPosition); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	task.Position = newPosition

	return nil
}

func (c *PostgresTasksStorage) Shift(task *models.Task, columnId int) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	var maxPosition int

	{
		queryMaxLeftPosition := `SELECT coalesce(max(position), -1) FROM tasks WHERE column_id = $1`
		stmt, err := tx.Prepare(queryMaxLeftPosition)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(columnId).Scan(&maxPosition); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	{
		sqlMove := `UPDATE tasks SET column_id = $2, position = position + $3 WHERE id = $1`
		stmt, err := tx.Prepare(sqlMove)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(task.Id, columnId, maxPosition+1); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	task.ColumnId = columnId

	return tx.Commit()
}

func (c *PostgresTasksStorage) Update(task *models.Task) error {
	sqlUpdate := `UPDATE tasks SET name = $2, description = $3 WHERE id = $1`
	_, err := postgres.DB().Exec(sqlUpdate, task.Id, task.Name, task.Description)
	return err
}

func (c *PostgresTasksStorage) Delete(task *models.Task) error {
	_, err := postgres.DB().Exec(`DELETE FROM tasks WHERE id = $1`, task.Id)
	return err
}

func (c *PostgresTasksStorage) CountInColumn(columnId int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM tasks WHERE column_id = $1`
	err := postgres.DB().QueryRow(query, columnId).Scan(&count)
	return count, err
}

func (c *PostgresTasksStorage) Clear() error {
	_, err := postgres.DB().Exec("TRUNCATE tasks RESTART IDENTITY CASCADE;")
	return err
}
