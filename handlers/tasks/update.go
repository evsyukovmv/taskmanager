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

func Update(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	tb := &models.TaskBase{}
	err = json.NewDecoder(r.Body).Decode(tb)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c, err := tasksvc.Update(taskId, tb)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
