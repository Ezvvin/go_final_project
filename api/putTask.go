package api

import (
	"encoding/json"
	"net/http"

	db "example/internal/db"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	if task.ID == "" {

		http.Error(w, "Не указан идентификатор задачи", http.StatusBadRequest)
		return
	}

	if err := task.CheckTask(); err != nil {
		http.Error(w, "Error"+err.Error(), http.StatusBadRequest)
		return
	}

	err = dbs.PutTask(task)
	if err != nil {
		http.Error(w, "Error:"+err.Error(), http.StatusInternalServerError)

		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{})
}
