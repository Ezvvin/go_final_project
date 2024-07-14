package main

import (
	"fmt"
	"log"
	"net/http"

	"example/api"
	"example/config"
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
	addr := fmt.Sprintf(":%s", config.Port)

	// Router
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir("./web")))

	r.Get("/api/nextdate", api.GetNextDateHandler)
	r.MethodFunc(http.MethodGet, "/api/task", api.GetTask)
	r.MethodFunc(http.MethodPut, "/api/task", api.UpdateTask)
	r.MethodFunc(http.MethodDelete, "/api/task", api.DeleteTask)
	r.MethodFunc(http.MethodPost, "/api/task", api.AddTask)
	r.MethodFunc(http.MethodGet, "/api/tasks", api.GetTasks)
	r.MethodFunc(http.MethodPost, "/api/task/done", api.TaskDone)

	log.Printf("Сервер на порту: %s\n", config.Port)
	// Запуск сервера
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Println(err)
	}
}
