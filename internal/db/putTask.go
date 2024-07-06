package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// PutTask отправляет SQL запрос на обновление задачи Task, возвращает ошибку в случае неудачи.
func (dbHandl *DB) PutTask(updateTask Task) error {
	res, err := dbHandl.db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", updateTask.Date),
		sql.Named("title", updateTask.Title),
		sql.Named("comment", updateTask.Comment),
		sql.Named("repeat", updateTask.Repeat),
		sql.Named("id", updateTask.ID))
	if err != nil {
		return err
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return fmt.Errorf("ошибка при обновление задачи")
	}
	return nil
}
