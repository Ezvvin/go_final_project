package database

import (
	"database/sql"
	"example/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// Получение пути для базы данных
func getDBFile() string {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Не удалось найти .env")
	}
	return os.Getenv("TODO_DBFILE")
}

// проверяем наличие файла БД
func checkDbFile() bool {
	_, err := os.Stat(getDBFile())
	//если БД-фаил не существует, возвращаем false
	if err != nil {
		fmt.Println("Database file is not created!")
		return false
	}
	//если БД-фаил создан, возвращаем true
	return true
}

// Открываем Базу данных
var db *sql.DB

func OpenDatabase() (*sql.DB, error) {
	if !checkDbFile() {
		//Создаем БД фаил, если его нет
		file, err := os.Create(getDBFile())
		if err != nil {
			return nil, err
		}
		log.Println("Created is database file")
		file.Close()
	}

	db, err := sql.Open("sqlite3", getDBFile())
	log.Println("Открытие базы данных")
	if err != nil {
		fmt.Println("Ошибка открытия базы данных.", err)
		return nil, err
	}
	defer db.Close()

	//Создаем таблицу в базе данных
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scheduler (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"date" CHAR(8) NOT NULL,
			"title" VARCHAR(255) NOT NULL,
			"comment" TEXT,
			"repeat" VARCHAR(128) NOT NULL
			)
		`)
	if err != nil {
		return nil, err
	}
	//Создаем индекс по дате
	_, err = db.Exec(
		`CREATE INDEX IF NOT EXISTS idx_scheduler_date ON sсheduler (date)`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Получение идентификатора созданной задачи
func InsertTask(task models.Task) (int, error) {
	result, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
