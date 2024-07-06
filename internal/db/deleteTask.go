package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DeleteTask отправялет SQL запрос на удаление задачи с указанным ID. Возваращает ошибку в случае неудачи.
func (dbHandl *DB) DeleteTask(id string) error {
	_, err := dbHandl.GetTaskByID(id)
	if err != nil {
		return err
	}

	res, err := dbHandl.db.Exec("DELETE FROM scheduler WHERE id= :id", sql.Named("id", id))
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected != 1 {
		return fmt.Errorf("при удаление что-то пошло не так")
	}
	return nil
}
