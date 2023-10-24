package module

import (
	"slot-server/internal/slot/model"
	"testing"
)

func TestCheckAllLinePay(t *testing.T) {
	var pay = []model.Line{
		{
			Symbol: 1,
			Payout: []float32{0, 0, 100, 200, 300},
		},
		{
			Symbol: 2,
			Payout: []float32{0, 0, 10, 20, 30},
		},
	}

	grid := [][]int{
		{21, 11, 21},
		{99, 21, 21},
		{99, 13, 12},
		{20, 10, 12},
		{12, 11, 22},
	}

	lp := AllLineWins(grid, pay, 1.0)

	if len(lp) != 2 {
		t.Fatalf("line pay count not matched | Want : %d , Res : %d", 2, len(lp))
	}

	if lp[0].Win != 2*2*3*1*2*300 {
		t.Fatalf("line pay win not matched | Want : %f , Res : %f\nMatched Line Count: %d, %d, %d, %d, %d", 2*2*3*1*2*300.0, lp[0].Win, lp[0].Position[0], lp[0].Position[1], lp[0].Position[2], lp[0].Position[3], lp[0].Position[4])
	}
}

func TestFindScatter(t *testing.T) {
	scatter := model.Scatter{
		Symbol:    0,
		Payout:    []float32{0, 0, 0},
		BonusType: model.FreeSpin,
		Bonus:     []int{0, 0, 15},
	}

	grid := [][]int{
		{1, 0, 2},
		{2, 1, 1},
		{1, 1, 0},
		{4, 99, 2},
		{5, 0, 1},
	}

	scWin := ScatterWin(grid, scatter, 10.0)

	if int32(scatter.BonusType) != scWin.Bonus || 15 != int(scWin.BonusParam) {
		t.Fatalf("scatter 3win error!!! %s", scWin)
	}

	grid = [][]int{
		{1, 1, 2},
		{2, 1, 1},
		{1, 1, 0},
		{4, 99, 2},
		{5, 0, 1},
	}

	scWin = ScatterWin(grid, scatter, 10.0)

	if scWin != nil {
		t.Fatalf("scatter 2win error!!! %s", scWin)
	}

	grid = [][]int{
		{1, 1, 2},
		{2, 1, 1},
		{1, 1, 5},
		{4, 99, 2},
		{5, 5, 1},
	}

	scWin = ScatterWin(grid, scatter, 10.0)

	if scWin != nil {
		t.Fatalf("scatter 0win error!!! %s", scWin)
	}
}
