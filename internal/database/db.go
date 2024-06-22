package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

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
		fmt.Println("БД не создана!")
		return false
	}
	//если БД-фаил создан, возвращаем true
	return true
}

// Открываем Базу данных
func OpenDatabase() (*sql.DB, error) {
	if !checkDbFile() {
		//Создаем БД
		file, err := os.Create(getDBFile())
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", getDBFile())
	fmt.Println("открытие ФАЙЛА БД")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

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
	_, err = db.Exec(
		`CREATE INDEX IF NOT EXISTS idx_date ON sсheduler (date)`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
