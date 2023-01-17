package main

import (
	"book-api/api"
	"book-api/db"
	"os"
)

func main() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:yCg0jfiFWZZBjFTXyU2I@containers-us-west-169.railway.app:6782/railway")
	db.Connect()
	api.Routes()
	api.Start()
}
