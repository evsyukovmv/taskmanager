package helpers

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/logger"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	logger.Info(err.Error())
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{ error: "` + err.Error() + `" }`))
}

func WriteJSON(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
