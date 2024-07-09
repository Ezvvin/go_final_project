package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// GetTaskByID возвращает задачу Task с указанным ID, или ошибку.
func (dbHandl *Storage) GetTaskByID(id string) (Task, error) {
	var task Task

	row := dbHandl.db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id))

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		log.Println(err)
		return Task{}, err
	}
	return task, nil

}
