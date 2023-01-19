package main

import (
	"book-api/api"
	"book-api/db"
)

func main() {
	db.Connect()
	api.Routes()
	api.Start()
}
