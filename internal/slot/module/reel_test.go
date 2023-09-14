package module

import (
	"testing"
)

func TestCheckAllLinePay(t *testing.T) {
	var pay = []Pay{
		{
			Symbol: 1,
			Payout: []float64{0, 0, 100, 200, 300},
		},
		{
			Symbol: 2,
			Payout: []float64{0, 0, 10, 20, 30},
		},
	}
	grid := [][]int{
		{1, 1, 2},
		{2, 1, 1},
		{1, 1, 1},
		{4, 99, 2},
		{5, 1, 1},
	}

	lp := AllLineWins(grid, pay, 1.0)

	if len(lp) != 2 {
		t.Fatalf("line pay count not matched | Want : %d , Res : %d", 2, len(lp))
	}

	if lp[0].Win != 2*2*3*1*2*300 {
		t.Fatalf("line pay win not matched | Want : %f , Res : %f\nMatched Line Count: %d, %d, %d, %d, %d", 2*2*3*1*2*300.0, lp[0].Win, lp[0].Position[0], lp[0].Position[1], lp[0].Position[2], lp[0].Position[3], lp[0].Position[4])
	}
}
