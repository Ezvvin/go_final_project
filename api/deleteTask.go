package api

import (
	"encoding/json"
	"net/http"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Не указан идентификатор задачи", http.StatusBadRequest)
		return
	}

	err := dbs.DeleteTask(id)
	if err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{})

}
