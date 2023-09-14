package slot00

import (
	"github.com/BurntSushi/toml"
	"log"
	"slot-server/internal/slot/model"
	"slot-server/internal/slot/module"
)

func Init() *ParSheet {
	var ps ParSheet
	_, err := toml.DecodeFile("./configs/slot00.toml", &ps)
	if err != nil {
		log.Fatal(err)
	}

	for _, strip := range ps.Strips {
		strip.StakingWeight()
	}

	return &ps
}

func (m *ParSheet) Spin(_ interface{}, bet float64) (*model.SpinOutput, *model.Error) {

	grid := module.GenRandomGrid(m.Strips, m.Column)

	lineWins := module.AllLineWins(grid, m.Pays, bet)

	totalWin := 0.0
	for _, win := range lineWins {
		totalWin += win.Win
	}

	upSymbols := []int{11, 12, 13, 14, 10}
	downSymbols := []int{11, 12, 13, 14, 10}

	bonus := Bonus{FreeSpin: "15"}

	return &model.SpinOutput{
		Win:         totalWin,
		TotalWin:    totalWin,
		Symbols:     grid,
		UpSymbols:   upSymbols,
		DownSymbols: downSymbols,
		LineWins:    lineWins,
		BonusWins:   []model.BonusWin{bonus},
		NextProcess: nil,
	}, nil
}
