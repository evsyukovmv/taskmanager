package tasks

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	response := fmt.Sprintf("TASKS: GetList projectId %q, columnId %q", projectId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Create(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	response := fmt.Sprintf("TASKS: Create projectId %q, columnId %q", projectId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func GetById(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	response := fmt.Sprintf("TASKS: Create GetById projectId %q, columnId %q, taskId %q", projectId, taskId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Update(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	response := fmt.Sprintf("TASKS: Update projectId %q, columnId %q, taskId %q", projectId, taskId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	response := fmt.Sprintf("TASKS: Delete projectId %q, columnId %q, taskId %q", projectId, taskId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
