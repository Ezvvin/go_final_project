package api

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	nd "example/internal/nextdate"
)

// PostTaskDoneHandler обрабатывает запросы к /api/task/done с методом POST.
// Если пользователь авторизован, удаляет задачи не имеющих правил повторения repeat, или обновляет дату выполнения задач, имеющих правило repeat.
// Возвращает пустой JSON {} в случае успеха, или JSON {"error": error} при возникновение ошибки.
func PostTaskDoneHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	q := r.URL.Query()
	id := q.Get("id")
	isID, _ := regexp.Match("[0-9]+", []byte(id))
	if !isID {
		writeErr(fmt.Errorf("некорректный формат id"), w)
		return
	}
	task, err := dbs.GetTaskByID(id)
	if err != nil {
		writeErr(err, w)
		return
	}
	if len(task.Repeat) == 0 {
		err = dbs.DeleteTask(id)
		if err != nil {
			writeErr(err, w)
			return
		}
		writeEmptyJson(w)
		return
	} else {
		nextDate, err := nd.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeErr(err, w)
			return
		}
		task.Date = nextDate
	}
	err = dbs.PutTask(task)
	if err != nil {
		writeErr(err, w)
		return
	}
	writeEmptyJson(w)

}
