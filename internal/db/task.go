package db

import (
	"fmt"
	"log"
	"regexp"
	"strings"
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

// formatTask проверяет переданную задачу Task на корректность полей, а так же корректирует дату задачи.
func (task Task) FormatTask() (Task, error) {
	var date time.Time
	var err error

	if len(task.Date) == 0 || strings.ToLower(task.Date) == "today" {
		date = time.Now()
		task.Date = date.Format(config.DateFormat)

	} else {
		date, err = time.Parse(config.DateFormat, task.Date)
		if err != nil {
			log.Println(err)
			return Task{}, err
		}
	}
	if isID, _ := regexp.Match("[0-9]+", []byte(task.ID)); !isID && task.ID != "" {
		err = fmt.Errorf("некорректный формат ID")
		return Task{}, err
	}

	// Даты с временем приведённым к 00:00:00
	dateTrunc := date.Truncate(time.Hour * 24)
	nowTrunc := time.Now().Truncate(time.Hour * 24)

	if dateTrunc.Before(nowTrunc) {
		switch {
		case len(task.Repeat) > 0:
			task.Date, err = nd.NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				log.Println(err)
				return Task{}, err
			}
		case len(task.Repeat) == 0:
			task.Date = time.Now().Format(config.DateFormat)
		}

	}
	return task, nil
}
