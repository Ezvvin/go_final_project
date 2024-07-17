package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"example/config"
	db "example/internal/db"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []db.Task
	var err error
	var date time.Time
	var resp []byte

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	q := r.URL.Query()
	search := q.Get("search")
	isDate, _ := regexp.Match("[0-9]{2}.[0-9]{2}.[0-9]{4}", []byte(search))

	switch {
	case len(search) == 0:
		tasks, err = dbs.GetTasksList()

	case isDate:
		date, err = time.Parse("02.01.2006", search)
		if err == nil {
			search = date.Format(config.DateFormat)
			tasks, err = dbs.GetTasksList(search)
			break
		}
		fallthrough

	default:
		search = fmt.Sprint("%" + search + "%")
		tasks, err = dbs.GetTasksList(search)

	}

	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return
	} else {
		if len(tasks) == 0 {
			tasksResp := map[string][]db.Task{
				"tasks": {},
			}
			resp, err = json.Marshal(tasksResp)
		} else {
			tasksResp := map[string][]db.Task{
				"tasks": tasks,
			}
			resp, err = json.Marshal(tasksResp)

		}

		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(resp)
		if err != nil {
			log.Println(err)
		}
		return
	}

}
