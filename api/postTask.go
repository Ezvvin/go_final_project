package api

import (
	"encoding/json"
	"log"
	"net/http"

	db "example/internal/db"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("Ошибка десериализации JSON:", err)
		http.Error(w, `{"error": "Ошибка десериализации JSON"}`, http.StatusBadRequest)
		return
	}

	id, err := dbs.AddTask(task)
	if err != nil {
		log.Println("Ошибка добавления задачи:", err)
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}
