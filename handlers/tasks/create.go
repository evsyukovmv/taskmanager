package tasks

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/tasksvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	columnId, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	t := &models.Task{ColumnId: columnId}

	err = json.NewDecoder(r.Body).Decode(&t.TaskBase)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	err = tasksvc.Create(t)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, t)
}
