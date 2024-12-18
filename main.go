package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// ============================== Env Variables
	// load env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// ============================== Router
	// Initialize router
	r := http.NewServeMux()

	// Routing

	// ============================== Server
	// Read server address from env
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Printf("app running on http://%s", addr)

	// Run server
	log.Fatal(http.ListenAndServe(addr, r))
}
