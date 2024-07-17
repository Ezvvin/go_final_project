package api

import (
	"encoding/json"
	"log"
	"net/http"

	db "example/internal/db"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("Ошибка десериализации JSON:", err)
		http.Error(w, `{"error": "Ошибка десериализации JSON"}`, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		log.Println("Не указан идентификатор задачи")
		http.Error(w, `{"error": "Не указан идентификатор задачи"}`, http.StatusBadRequest)
		return
	}

	if err := task.CheckTask(&task); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	err = dbs.PutTask(task)
	if err != nil {
		if err.Error() == "задача не найдена" {
			http.Error(w, `{"error": "Задача не найдена"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{})
}
