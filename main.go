package main

import (
	"github.com/evsyukovmv/taskmanager/handlers"
	"github.com/evsyukovmv/taskmanager/logger"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/evsyukovmv/taskmanager/services/columnsvc"
	"github.com/evsyukovmv/taskmanager/services/projectsvc"
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
	err := postgres.Configure(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Postgres failed: %s", err.Error())
	}
}

func configureServices() {
	projectsvc.NewService(&pgstorage.PostgresProjectsStorage{})
	columnsvc.NewService(&pgstorage.PostgresColumnsStorage{})
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
