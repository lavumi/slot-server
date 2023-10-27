package main

import (
	"slot-server/internal/server"
)

func main() {
	//
	//err := godotenv.Load(".web.dev.env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	server.Run()

}
