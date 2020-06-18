package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/columnsvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c := &models.Column{ProjectId: projectId}

	err = json.NewDecoder(r.Body).Decode(&c.ColumnBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = columnsvc.Create(c)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
