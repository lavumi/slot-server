package main

import (
	"github.com/joho/godotenv"
	"log"
	"slot-server/internal/server"
	"slot-server/internal/slot"
)

func main() {

	err := godotenv.Load(".web.dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	go slot.Run()
	server.Run()

}
