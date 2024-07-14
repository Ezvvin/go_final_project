package api

import (
	"encoding/json"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	search := r.URL.Query().Get("search")
	tasks, err := dbs.GetTasksList(search)
	if err != nil {
		http.Error(w, "Ошибка выполнения запроса списка.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"tasks": tasks})

}
