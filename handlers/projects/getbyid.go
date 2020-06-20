package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/projectsvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	p, err := projectsvc.GetById(id)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, p)
}
