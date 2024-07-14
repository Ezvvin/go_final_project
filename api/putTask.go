package api

import (
	"encoding/json"
	"errors"
	"example/config"
	db "example/internal/db"
	nd "example/internal/nextdate"
	"net/http"
	"time"
)

var task db.Task

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	if task.ID == "" {

		http.Error(w, "Не указан идентификатор задачи", http.StatusBadRequest)
		return
	}

	if err := сheckTask(&task); err != nil {
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

func сheckTask(task *db.Task) error {
	if task.Title == "" {
		return errors.New("не указан заголовок задачи")
	}

	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(config.DateFormat)
	} else {
		date, err := time.Parse(config.DateFormat, task.Date)
		if err != nil {
			return errors.New("дата представлена в неправильном формате")
		}

		if date.Before(now) {
			if task.Repeat == "" {
				task.Date = now.Format(config.DateFormat)
			} else {
				nextDate, err := nd.NextDate(now, task.Date, task.Repeat)
				if err != nil {
					return errors.New("ошибка вычисления следующей даты")
				}
				task.Date = nextDate
			}
		}
	}

	return nil
}
