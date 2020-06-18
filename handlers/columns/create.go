package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/columns"
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

	column := &models.Column{ProjectId: projectId}

	err = json.NewDecoder(r.Body).Decode(&column.ColumnBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(column)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = columns.Storage().Create(column)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, column)
}
