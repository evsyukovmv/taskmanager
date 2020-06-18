package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/columns"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
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

	err = json.NewDecoder(r.Body).Decode(&c.ColumnBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(c)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = columns.Storage().Update(c)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
