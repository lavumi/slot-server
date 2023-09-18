package main

import (
	"slot-server/internal/server"
	"slot-server/internal/slot"
)

func main() {

	go slot.Run()
	server.Run()

}
