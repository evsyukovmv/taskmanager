package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/postgres"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	column, err := findColumn(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&column.ColumnBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(column)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = postgres.DB().Update(column)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, column)
}
