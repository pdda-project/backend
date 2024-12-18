package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pdda-project/backend/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// ============================== Env Variables
	// load env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// db
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_URI")), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// ============================== Router
	// Initialize router
	r := http.NewServeMux()

	// Routing
	userHandler := users.NewUserHandler(db)
	userHandler.RegisterRouter(r)

	// ============================== Server
	// Read server address from env
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Printf("app running on http://%s", addr)

	// Run server
	log.Fatal(http.ListenAndServe(addr, r))
}
