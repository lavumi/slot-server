package main

import "slot-server/internal/slot"

func main() {
	connect, err := slot.Connect()
	if err != nil {
		return
	}

	_, _, err = connect.RequestSpin(0, 1.0, "")
	if err != nil {
		return
	}

}
