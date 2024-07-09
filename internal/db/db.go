package db

import (
	"database/sql"
	"example/config"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

// CheckDbFile проверяет существует ли файл переданный аргументом
func CheckDbFile(dbFile string) bool {
	_, err := os.Stat(dbFile)
	var check bool
	if err == nil {
		check = true
	}
	return check

}

// OpenDatabase открывает базу данных указанную в .env файле, добавляет её в структуру DBHandler.
func OpenDatabase() (Storage, error) {
	db, err := sql.Open("sqlite3", config.DbFile)
	if err != nil {
		return Storage{}, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(6)
	db.SetConnMaxIdleTime(time.Minute * 5)
	dbHandl := Storage{}
	dbHandl.db = db

	return dbHandl, nil
}

// CloseDB закрывает подключение к базе данных.
func (dbHandl *Storage) CloseDB() error {
	db := dbHandl.db
	err := db.Close()
	if err != nil {
		return err
	}
	return nil

}

// InstallDB создаёт файл для базы данных с названием, указаным в .env,
// отправляет SQL запрос на создание таблицы из файла qu.sql.
// Возвращает ошибку в случае неудачи.
func InstallDB() error {
	db, err := sql.Open("sqlite3", config.DbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	createTable := `CREATE TABLE "scheduler" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"date"	TEXT NOT NULL,
	"title"	TEXT NOT NULL,
	"comment"	TEXT,
	"repeat"	TEXT NOT NULL DEFAULT "",
	CHECK(length("repeat") <= 128)
	CHECK(length("title") > 0)
);`

	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	// Создаем индекс по полю date для сортировки задач по дате
	createIndexDate := `CREATE INDEX "scheduler_date" ON "scheduler" (
	"date"	DESC
);`
	_, err = db.Exec(createIndexDate)
	if err != nil {
		return err
	}

	return nil
}
