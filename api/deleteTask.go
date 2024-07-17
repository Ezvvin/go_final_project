package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error": "Не указан идентификатор задачи"}`, http.StatusBadRequest)
		return
	}

	err := dbs.DeleteTask(id)
	if err != nil {
		if err.Error() == "задача не найдена" {
			log.Println("Задача не найдена", err)
			http.Error(w, `{"error": "Задача не найдена"}`, http.StatusNotFound)
		} else {
			log.Println("Ошибка удаления задачи", err)
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{})
}
