package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Dotenv configs
	host := os.Getenv("USER_SERVICE_HOST")
	dbUri := os.Getenv("DB_URI")

	// Create database
	db, err := gorm.Open(sqlite.Open(dbUri))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create server
	server, err := NewGRPCServer(host, db)
	if err != nil {
		log.Fatal(err.Error())
	}
	server.Run()
}
