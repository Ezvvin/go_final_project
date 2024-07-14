package api

import (
	"log"
	"net/http"
	"time"

	"example/config"
	db "example/internal/db"
	nd "example/internal/nextdate"
)

var (
	dbs db.Storage
)

// InitApi инициплизирует переменные используемые в пакете api, зависящие от переменных среды и других пакетов
func InitApi(storage db.Storage) {
	dbs = storage
}

// GetNextDateHandler обрабатывает GET запросы к api/nextdate
func GetNextDateHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	now := q.Get("now")
	date := q.Get("date")
	repeat := q.Get("repeat")
	nowDate, err := time.Parse(config.DateFormat, now)
	if err != nil {
		log.Println(err)
		return
	}

	nextDate, err := nd.NextDate(nowDate, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := []byte(nextDate)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}
