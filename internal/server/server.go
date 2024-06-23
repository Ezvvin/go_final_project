package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"example/internal/transport/rest"
)

// Получение порта локально
func getPort() string {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Не удалось найти .env")
	}
	return ":" + os.Getenv("TODO_PORT")
}

// запуск сервера
func StartServer() {
	const webDir = "../../web"
	handler := chi.NewRouter()
	handler.Mount("/", http.FileServer(http.Dir(webDir)))
	fmt.Println(handler)
	// Обработчик API для вычисления следующей даты
	handler.Get("/api/nextdate", rest.HandleNextDate)
	// Обработчик API для добавления задачи
	handler.Post("/api/task", rest.TaskAddPOST)

	log.Println("Сервер запущен на порту", getPort())
	if err := http.ListenAndServe(getPort(), handler); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
