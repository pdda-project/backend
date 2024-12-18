package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pdda-project/backend/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Running migrations...")

	// Include models here
	err = db.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Migration success")

}
