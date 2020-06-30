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

func Update(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	cb := &models.CommentBase{}
	err = json.NewDecoder(r.Body).Decode(cb)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	c, err := services.ForComment().Update(commentId, cb)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, c)
}
