package tasks

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/tasksvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c, err := tasksvc.GetById(taskId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
