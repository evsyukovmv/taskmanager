package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	pb := &models.ProjectBase{}
	err = json.NewDecoder(r.Body).Decode(pb)
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	p, err := services.ForProject().Update(id, pb)
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	helpers.WriteJSON(w, r, p)
}
