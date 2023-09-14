package main

import (
	"fmt"
	"slot-server/internal/slot"
	"slot-server/internal/slot/model"
)

func main() {
	m := slot.Initialize()

	testReq := model.SpinInput{
		Id:        0,
		BetCash:   0.5,
		BetLine:   20,
		PrevState: nil,
	}

	r, err := m.Spin(testReq)
	if err != nil || r == nil {
		fmt.Printf("%s", err.Error())
		return
	}

	fmt.Printf("Win : %f\n", r.Win)
	fmt.Printf("TotalWin : %f\n", r.TotalWin)
	fmt.Printf("Symbols : %v\n", r.Symbols)
	fmt.Printf("UpSymbols : %v\n", r.UpSymbols)
	fmt.Printf("DownSymbols : %v\n", r.DownSymbols)
	fmt.Printf("LineWins : %v\n", r.LineWins)
}
