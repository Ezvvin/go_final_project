package db

import (
	"errors"
	"time"

	"example/config"
	nd "example/internal/nextdate"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (task *Task) CheckTask(t *Task) error {
	if t.Title == "" {
		return errors.New("не указан заголовок задачи")
	}

	now := time.Now()
	if t.Date == "" {
		t.Date = now.Format(config.DateFormat)
	} else {
		date, err := time.Parse(config.DateFormat, t.Date)
		if err != nil {
			return errors.New("дата представлена в неправильном формате")
		}

		if date.Before(now) {
			if t.Repeat == "" {
				t.Date = now.Format(config.DateFormat)
			} else {
				nextDate, err := nd.NextDate(now, t.Date, t.Repeat)
				if err != nil {
					return errors.New("ошибка вычисления следующей даты")
				}
				t.Date = nextDate
			}
		}
	}

	return nil
}
