package services

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/logger"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/evsyukovmv/taskmanager/storages/pgstorage"
	"os"
)

func setupTests() error {
	if err := logger.Configure(); err != nil {
		return fmt.Errorf("logger failed: %s", err.Error())
	}
	if err := postgres.Configure(os.Getenv("DATABASE_URL")); err != nil {
		return fmt.Errorf("postgres failed: %s", err.Error())
	}
	NewProjectService(&pgstorage.PostgresProjectsStorage{})
	NewColumnService(&pgstorage.PostgresColumnsStorage{})
	NewTaskService(&pgstorage.PostgresTasksStorage{})
	NewCommentService(&pgstorage.PostgresCommentsStorage{})
	return nil
}

func clearTests() {
	_ = ForComment().Clear()
	_ = ForTask().Clear()
	_ = ForColumn().Clear()
	_ = ForProject().Clear()
}
