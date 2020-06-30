package tasks

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Move(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	tp := &models.TaskPosition{}
	err = json.NewDecoder(r.Body).Decode(tp)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	c, err := services.ForTask().Move(taskId, tp)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}


	helpers.WriteJSON(w, r, c)
}
