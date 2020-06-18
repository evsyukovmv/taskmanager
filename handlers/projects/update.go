package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/projects"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	p, err := projects.Storage().GetByID(id)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&p.ProjectBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(p)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	projects.Storage().Update(p)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, p)
}

