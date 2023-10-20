package main

import (
	"fmt"
	"slot-server/internal/slot"
)

func main() {
	connect, err := slot.Connect()
	if err != nil {
		return
	}

	spin, state, cash, err := connect.RequestSpin(0, 1.0, nil)
	if err != nil {
		return
	}

	fmt.Println(string(spin))
	fmt.Println(string(state))
	fmt.Println(cash)
}
