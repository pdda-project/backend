package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pdda-project/backend/services/users/repo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbUri := os.Getenv("DB_URI")
	db, err := gorm.Open(sqlite.Open(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&repo.User{})
}
