package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	host := os.Getenv("USER_SERVICE_HOST")

	server, err := NewGRPCServer(host)
	if err != nil {
		log.Fatal(err.Error())
	}

	server.Run()

	// Tangani sinyal untuk terminasi server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	server.Close()
}
