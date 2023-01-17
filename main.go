package main

import (
	"book-api/api"
	"book-api/db"
	"os"
)

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:asd123@localhost:5432/book-api")
	db.Connect()
	api.Routes()
	api.Start()
}
