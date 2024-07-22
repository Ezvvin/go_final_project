package api

import (
	"log"
	"net/http"
	"time"

	"example/config"
	nd "example/internal/nextdate"
)

// GetNextDate обрабатывает GET запросы к api/nextdate
func GetNextDate(w http.ResponseWriter, r *http.Request) {

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
		log.Println(err)
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
