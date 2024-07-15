package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// GetTaskByID возвращает задачу Task с указанным ID, или ошибку.
func (dbHandl *Storage) GetTaskByID(id string) (Task, error) {
	var task Task

	row := dbHandl.db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id))

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		log.Println("Ошибка выполнения запроса:", err)
		return task, err
	}
	return task, nil

}
