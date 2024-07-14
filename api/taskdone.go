package api

import (
	"encoding/json"
	nd "example/internal/nextdate"
	"net/http"
	"time"
)

func TaskDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Не указан идентификатор задачи", http.StatusBadRequest)
		return
	}

	task, err := dbs.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	if task.Repeat == "" {
		// Удаляем одноразовую задачу
		err := dbs.DeleteTask(id)
		if err != nil {
			http.Error(w, "Error:"+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Рассчитываем следующую дату для периодической задачи
		now := time.Now()
		nextDate, err := nd.NextDate(now, task.Date, task.Repeat)
		if err != nil {

			http.Error(w, "Ошибка вычисления следующей даты", http.StatusInternalServerError)
			return
		}

		// Обновляем задачу с новой датой
		task.Date = nextDate
		err = dbs.PutTask(task)
		if err != nil {
			http.Error(w, "Error:"+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{})
}
