package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	RowsLimit = 15
)

var (
	Port       string
	DbFile     string
	DateFormat string
)

func EnvLoad() error {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		return err
	}

	Port = os.Getenv("TODO_PORT")
	DbFile = os.Getenv("TODO_DBFILE")
	DateFormat = os.Getenv("TODO_DATEFORMAT")
	return nil
}
