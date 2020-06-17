package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/postgres"
	"net/http"
)

func Move(w http.ResponseWriter, r *http.Request) {
	column, err := findColumn(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	oldPosition := column.Position
	err = json.NewDecoder(r.Body).Decode(&column.ColumnPosition)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	if oldPosition == column.Position {
		helpers.WriteJSON(w, column)
		return
	}

	tx, err := postgres.DB().Begin()
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	_, err = postgres.DB().Exec("UPDATE columns SET position = NULL WHERE id = ?", column.Id)
	if err != nil {
		_ = tx.Rollback()
		helpers.WriteError(w, err)
		return
	}

	_, err = postgres.DB().Exec(
		"UPDATE columns SET position = position + 1 WHERE project_id = ?",
		column.ProjectId, column.Position, column.Id,
	)
	if err != nil {
		_ = tx.Rollback()
		helpers.WriteError(w, err)
		return
	}

	_, err = postgres.DB().Exec("UPDATE columns SET position = ? WHERE id = ?", column.Position, column.Id)
	if err != nil {
		_ = tx.Rollback()
		helpers.WriteError(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, column)
}
