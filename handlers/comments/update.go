package comments

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/commentsvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	cb := &models.CommentBase{}
	err = json.NewDecoder(r.Body).Decode(cb)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c, err := commentsvc.Update(commentId, cb)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
