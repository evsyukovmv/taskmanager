package pgstorage

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
)

type PostgresProjectsStorage struct {}

func (p *PostgresProjectsStorage) GetList() (*[]models.Project, error) {
	var projects []models.Project

	rows, err := postgres.DB().Query("SELECT * FROM projects ORDER BY name ASC")
	if err != nil {
		return &projects, err
	}

	for rows.Next() {
		var project models.Project

		if err := rows.Scan(&project.Id, &project.Name, &project.Description); err != nil {
			return &projects, err
		}

		projects = append(projects, project)
	}
	return &projects, err
}

func (p *PostgresProjectsStorage) GetById(id int) (*models.Project, error) {
	project := &models.Project{Id: id}
	err := postgres.DB().QueryRow(
		"SELECT * FROM projects WHERE id = $1", project.Id).Scan(
		&project.Id, &project.Name, &project.Description)
	return project, err
}

func (p *PostgresProjectsStorage) Create(project *models.Project) error {
	tx, err := postgres.DB().Begin()
	if err != nil {
		return err
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO projects (name, description) VALUES ($1, $2) RETURNING id`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if err := stmt.QueryRow(project.Name, project.Description).Scan(&project.Id); err != nil {
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

		if _, err := stmt.Exec("default", project.Id); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (p *PostgresProjectsStorage) Update(project *models.Project) error {
	sqlUpdate := `UPDATE projects SET name = $2, description = $3 WHERE id = $1`
	_, err := postgres.DB().Exec(sqlUpdate, project.Id, project.Name, project.Description)
	return err
}

func (p *PostgresProjectsStorage) Delete(project *models.Project) error {
	sqlDelete := `DELETE FROM projects WHERE id = $1;`
	_, err := postgres.DB().Exec(sqlDelete, project.Id)
	return err
}
