package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/columns"
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

	c, err := columns.Storage().GetByID(columnId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	oldPosition := c.Position
	err = json.NewDecoder(r.Body).Decode(&c.ColumnPosition)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	if oldPosition == c.Position {
		helpers.WriteJSON(w, c)
		return
	}

	err = columns.Storage().Move(c)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
