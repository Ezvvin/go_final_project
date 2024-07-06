package db

import (
	"database/sql"
)

// AddTask отправляет SQL запрос на добавление переданной задачи Task. Возвращает ID добавленной задачи и/или ошибку.
func (dbHandl *DB) AddTask(task Task) (int64, error) {
	var id int64
	res, err := dbHandl.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date), sql.Named("title", task.Title),
		sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))
	if err == nil {
		id, _ = res.LastInsertId()
	}
	return id, err
}
