package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func getPort() string {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Не удалось найти .env")
	}
	return ":" + os.Getenv("TODO_PORT")
}
func main() {
	const webDir = "./web"
	handler := chi.NewRouter()
	handler.Mount("/", http.FileServer(http.Dir(webDir)))

	fmt.Printf("Starting server on port %s\n", getPort())
	if err := http.ListenAndServe(getPort(), handler); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
