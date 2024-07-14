package api

import (
	"encoding/json"
	"net/http"
)

func GetTask(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	task, err := dbs.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)

}
