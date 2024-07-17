package db

import (
	"errors"
	"time"

	"example/config"
	nd "example/internal/nextdate"
)

// AddTask отправляет SQL запрос на добавление переданной задачи Task. Возвращает ID добавленной задачи и/или ошибку.
func (dbHandl *Storage) AddTask(task Task) (int64, error) {
	if task.Title == "" {
		return 0, errors.New("пустой заголовок")
	}
	if task.Date == "" {
		task.Date = time.Now().Format(config.DateFormat)
	} else {
		_, err := time.Parse(config.DateFormat, task.Date)
		if err != nil {
			return 0, errors.New("неверный формат даты")
		}
	}

	now := time.Now()
	if task.Date < now.Format(config.DateFormat) {
		if task.Repeat == "" {
			task.Date = now.Format(config.DateFormat)
		} else {
			nextDate, err := nd.NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return 0, err
			}
			task.Date = nextDate
		}
	}

	result, err := dbHandl.db.Exec(
		`INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`,
		task.Date, task.Title, task.Comment, task.Repeat,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
