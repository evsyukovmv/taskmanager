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

func Shift(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	tc := &models.TaskColumn{}
	err = json.NewDecoder(r.Body).Decode(tc)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	c, err := tasksvc.Shift(taskId, tc)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}


	helpers.WriteJSON(w, r, c)
}