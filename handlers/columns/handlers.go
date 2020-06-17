package columns

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	response := fmt.Sprintf("COLUMNS: Create projectId %q", projectId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Update(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	response := fmt.Sprintf("COLUMNS: Update projectId %q, columnId %q", projectId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	columnId := chi.URLParam(r, "columnId")
	response := fmt.Sprintf("COLUMNS: Delete projectId %q, columnId %q", projectId, columnId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
