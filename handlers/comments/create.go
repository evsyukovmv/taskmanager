package comments

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	c := &models.Comment{TaskId: taskId}

	err = json.NewDecoder(r.Body).Decode(&c.CommentBase)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	err = services.ForComment().Create(c)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, c)
}
