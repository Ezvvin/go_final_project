package main

import (
	"fmt"
	"log"
	"net/http"

	"example/api"
	"example/config"
	"example/internal/auth"
	db "example/internal/db"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Загружаем переменные среды
	config.EnvLoad()
	// Если бд не существует, создаём
	if !db.CheckDbFile(config.DbFile) {
		err := db.InstallDB()
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
	addr := fmt.Sprintf("%s:%s", ip, config.Port)

	// Router
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir("./web")))

	r.Get("/api/nextdate", api.GetNextDateHandler)
	r.Get("/api/tasks", auth.Auth(api.GetTasksHandler))
	r.Post("/api/task/done", auth.Auth(api.PostTaskDoneHandler))
	r.Post("/api/signin", auth.Auth(api.PostSigninHandler))
	r.Handle("/api/task", auth.Auth(api.TaskHandler))

	log.Printf("Server running on %s\n", config.Port)
	// Запуск сервера
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Println(err)
	}
}
