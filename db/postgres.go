package db

import (
	"fmt"
	"log"
	"os"

	model "book-api/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL2"),
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&model.Book{}, &model.User{})

	DB = db

}
