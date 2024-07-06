package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

var (
	DateFormat string
)

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
func OpenDatabase() (DB, error) {
	dbFile := os.Getenv("TODO_DBFILE")
	DateFormat = os.Getenv("TODO_DATEFORMAT")

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return DB{}, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(6)
	db.SetConnMaxIdleTime(time.Minute * 5)
	dbHandl := DB{}
	dbHandl.db = db

	return dbHandl, nil
}

// CloseDB закрывает подключение к базе данных.
func (dbHandl *DB) CloseDB() error {
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
	dbFile := os.Getenv("TODO_DBFILE")
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	createTable := `CREATE TABLE "scheduler" (
	"id"	INTEGER,
	"date"	TEXT NOT NULL,
	"title"	TEXT NOT NULL,
	"comment"	TEXT,
	"repeat"	TEXT NOT NULL DEFAULT "",
	CHECK(length("repeat") <= 128)
	CHECK(length("title") > 0)
	PRIMARY KEY("id" AUTOINCREMENT)
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
