package main

import (
	"github.com/evsyukovmv/taskmanager/handlers"
	"github.com/evsyukovmv/taskmanager/logger"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/evsyukovmv/taskmanager/services/projects"
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
	err := postgres.Configure(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	if err != nil {
		log.Fatalf("Postgres failed: %s", err.Error())
	}
}

func configureServices() {
	projects.NewService(&pgstorage.PostgresProjectsStorage{})
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
