package main

import (
	"github.com/evsyukovmv/taskmanager/handlers"
	"github.com/evsyukovmv/taskmanager/logger"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/evsyukovmv/taskmanager/services"
	"github.com/evsyukovmv/taskmanager/storages/pgstorage"
	"log"
	"net/http"
	"os"
)

func main() {
	configureLogger()
	configurePostgres()
	configureServices()
	startHTTPServer()
}

func configureLogger() {
	err := logger.Configure()
	if err != nil {
		log.Fatalf("Logger failed: %s", err.Error())
	}
}

func configurePostgres() {
	err := postgres.Configure(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Postgres failed: %s", err.Error())
	}
}

func configureServices() {
	services.NewProjectService(&pgstorage.PostgresProjectsStorage{})
	services.NewColumnService(&pgstorage.PostgresColumnsStorage{})
	services.NewTaskService(&pgstorage.PostgresTasksStorage{})
	services.NewCommentService(&pgstorage.PostgresCommentsStorage{})
}

func startHTTPServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlers.NewRouter(),
	}
	logger.Info("Listening requests on :" + port + "...")
	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal("ListenAndServe failed on :" + port)
	}
}
