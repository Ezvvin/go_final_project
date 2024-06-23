package main

import (
	db "example/internal/database"
	"example/internal/server"
)

func main() {
	db.OpenDatabase()
	server.StartServer()
}
