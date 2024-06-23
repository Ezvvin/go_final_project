package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"example/internal/database"
	"example/internal/models"
	"example/internal/services"
	"fmt"
	"log"
	"net/http"
	"time"
)

func responseWithError(w http.ResponseWriter, errorText string, err error) {
	errorResponse := models.ErrorResponse{
		Error: fmt.Errorf("%s: %w", errorText, err).Error()}
	errorData, _ := json.Marshal(errorResponse)
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(errorData)

	if err != nil {
		http.Error(w, fmt.Errorf("error: %w", err).Error(), http.StatusBadRequest)
	}
}

func HandleNextDate(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeat := r.FormValue("repeat")

	now, err := time.Parse(models.DatePattern, nowStr)
	if err != nil {
		http.Error(w, "Invalid now format", http.StatusBadRequest)
		return
	}

	nextDate, err := services.NextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(nextDate))
}

func TaskAddPOST(w http.ResponseWriter, r *http.Request) {
	var taskData models.Task
	var buffer bytes.Buffer

	if _, err := buffer.ReadFrom(r.Body); err != nil {
		responseWithError(w, "body getting error", err)
		return
	}

	if err := json.Unmarshal(buffer.Bytes(), &taskData); err != nil {
		responseWithError(w, "JSON encoding error", err)
		return
	}

	if len(taskData.Date) == 0 {
		taskData.Date = time.Now().Format(models.DatePattern)
	} else {
		date, err := time.Parse(models.DatePattern, taskData.Date)
		if err != nil {
			responseWithError(w, "bad data format", err)
			return
		}

		if date.Before(time.Now()) {
			taskData.Date = time.Now().Format(models.DatePattern)
		}
	}

	if len(taskData.Title) == 0 {
		responseWithError(w, "invalid title", errors.New("title is empty"))
		return
	}

	if len(taskData.Repeat) > 0 {
		if _, err := services.NextDate(time.Now(), taskData.Date, taskData.Repeat); err != nil {
			responseWithError(w, "invalid repeat format", errors.New("no such format"))
			return
		}
	}

	taskId, err := database.InsertTask(taskData)
	if err != nil {
		responseWithError(w, "failed to create task", err)
		return
	}

	taskIdData, _ := json.Marshal(models.TaskIdResp{Id: taskId})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(taskIdData)
	log.Println("Added task with id=", taskId)

	if err != nil {
		responseWithError(w, "writing task id error", err)
	}
}
