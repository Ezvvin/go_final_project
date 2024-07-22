package api

import (
	db "example/internal/db"
)

var (
	dbs db.Storage
)

// InitApi инициплизирует переменные используемые в пакете api, зависящие от переменных среды и других пакетов
func InitApi(storage db.Storage) {
	dbs = storage
}
