package comments

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	response := fmt.Sprintf("COMMENTS: GetList projectId %q, columnId %q, taskId %q", projectId, columnId, taskId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Create(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	response := fmt.Sprintf("COMMENTS: Create projectId %q, columnId %q, taskId %q", projectId, columnId, taskId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func GetById(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	commentId := chi.URLParam(r, "commentId")
	response := fmt.Sprintf("COMMENTS: Create GetById projectId %q, columnId %q, taskId %q, commendId %q", projectId, taskId, columnId, commentId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Update(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	commentId := chi.URLParam(r, "commentId")
	response := fmt.Sprintf("COMMENTS: Update projectId %q, columnId %q, taskId %q, commendId %q", projectId, taskId, columnId, commentId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	taskId := chi.URLParam(r, "taskId")
	commentId := chi.URLParam(r, "commentId")
	response := fmt.Sprintf("COMMENTS: Delete projectId %q, columnId %q, taskId %q, commendId %q", projectId, taskId, columnId, commentId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
