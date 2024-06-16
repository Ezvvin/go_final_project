package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func getDBFile() string {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Не удалось найти .env")
	}
	return ":" + os.Getenv("TODO_DBFile")
}

func CreateDatabase() {
	db, err := sql.Open("sqlite", getDBFile())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}
