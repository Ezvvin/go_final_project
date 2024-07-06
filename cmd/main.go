package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example/api"
	"example/internal/auth"
	db "example/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Загружаем переменные среды
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	dbFile := os.Getenv("TODO_DBFILE")

	// Если бд не существует, создаём
	if !db.CheckDbFile(dbFile) {
		err = db.InstallDB()
		if err != nil {
			log.Println(err)
		}
	}

	// Запуск бд
	dbHandl, err := db.OpenDatabase()
	defer dbHandl.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	api.InitApi(dbHandl)

	// Адрес для запуска сервера
	ip := ""
	port := os.Getenv("TODO_PORT")
	addr := fmt.Sprintf("%s:%s", ip, port)

	// Router
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir("../web")))

	r.Get("/api/nextdate", api.GetNextDateHandler)
	r.Get("/api/tasks", auth.Auth(api.GetTasksHandler))
	r.Post("/api/task/done", auth.Auth(api.PostTaskDoneHandler))
	r.Post("/api/signin", auth.Auth(api.PostSigninHandler))
	r.Handle("/api/task", auth.Auth(api.TaskHandler))

	log.Printf("Server running on %s\n", port)
	// Запуск сервера
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Println(err)
	}
}
