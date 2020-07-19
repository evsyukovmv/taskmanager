package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
)

type PostgresCommentsStorage struct{}

func (c *PostgresCommentsStorage) GetListByTaskId(taskId int) (*[]models.Comment, error) {
	var comments []models.Comment

	rows, err := postgres.DB().Query("SELECT * FROM comments WHERE task_id = $1 ORDER BY created_at DESC", taskId)
	if err != nil {
		return &comments, err
	}

	for rows.Next() {
		var comment models.Comment

		if err := rows.Scan(&comment.Id, &comment.Text, &comment.CreatedAt, &comment.TaskId); err != nil {
			return &comments, err
		}

		comments = append(comments, comment)
	}

	return &comments, err
}

func (c *PostgresCommentsStorage) GetById(id int) (*models.Comment, error) {
	comment := &models.Comment{Id: id}
	err := postgres.DB().QueryRow(
		"SELECT * FROM comments WHERE id = $1", comment.Id).Scan(
		&comment.Id, &comment.Text, &comment.CreatedAt, &comment.TaskId)
	return comment, err
}

func (c *PostgresCommentsStorage) Create(comment *models.Comment) error {
	insertQuery := `INSERT INTO comments (text, task_id) VALUES ($1, $2) RETURNING id, created_at`
	return postgres.DB().QueryRow(insertQuery, comment.Text, comment.TaskId).Scan(&comment.Id, &comment.CreatedAt)
}

func (c *PostgresCommentsStorage) Update(comment *models.Comment) error {
	sqlUpdate := `UPDATE comments SET text = $2 WHERE id = $1`
	_, err := postgres.DB().Exec(sqlUpdate, comment.Id, comment.Text)
	return err
}

func (c *PostgresCommentsStorage) Delete(comment *models.Comment) error {
	_, err := postgres.DB().Exec(`DELETE FROM comments WHERE id = $1`, comment.Id)
	return err
}

func (c *PostgresCommentsStorage) Clear() error {
	_, err := postgres.DB().Exec("TRUNCATE comments RESTART IDENTITY CASCADE;")
	return err
}
