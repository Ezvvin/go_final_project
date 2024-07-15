package db

import (
	"database/sql"
	"log"
	"time"

	"example/config"
)

// GetTasksList возвращает последние добавленные задачи []Task, либо последние добавленные задачи подходящие под поисковой запрос search при его наличие.
func (dbHandl *Storage) GetTasksList(search ...string) ([]Task, error) {
	var tasks []Task
	var rows *sql.Rows
	var err error

	switch {
	case len(search) == 0:
		rows, err = dbHandl.db.Query("SELECT * FROM scheduler ORDER BY id LIMIT :limit", sql.Named("limit", config.RowsLimit))
		if err != nil {
			return []Task{}, err
		}
	case len(search) > 0:
		search := search[0]
		_, err = time.Parse(config.DateFormat, search)
		if err != nil {
			rows, err = dbHandl.db.Query("SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit",
				sql.Named("search", search),
				sql.Named("limit", config.RowsLimit))
			if err != nil {
				return []Task{}, err
			}
			break
		}
		rows, err = dbHandl.db.Query("SELECT * FROM scheduler WHERE date = :date LIMIT :limit",
			sql.Named("date", search),
			sql.Named("limit", config.RowsLimit))
		if err != nil {
			return []Task{}, err
		}
	}

	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		if err = rows.Err(); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return tasks, nil
}
