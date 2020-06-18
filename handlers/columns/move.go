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

func Move(w http.ResponseWriter, r *http.Request) {
	columnId, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	cp := &models.ColumnPosition{}
	err = json.NewDecoder(r.Body).Decode(cp)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c, err := columnsvc.Move(columnId, cp)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}


	helpers.WriteJSON(w, c)
}
