package foodie

import (
	"github.com/BurntSushi/toml"
	"log"
	"slot-server/internal/slot/api"
	"slot-server/internal/slot/api/proto"
	"slot-server/internal/slot/model"
	"slot-server/internal/slot/module"
)

var ps ParSheet

func init() {

	_, err := toml.DecodeFile("./parSheet/foodie.toml", &ps)
	if err != nil {
		log.Fatal(err)
	}

	for _, strip := range ps.Strips {
		strip.StakingWeight()
	}
}

func Spin(req *proto.Request) (*proto.Response, *model.Error) {
	grid := module.GenRandomGrid(ps.Strips, ps.Column)
	lineWins := module.AllLineWins(grid, ps.Pays, float64(req.BetCash))

	totalWin := 0.0
	for _, win := range lineWins {
		totalWin += win.Win
	}

	upSymbols := []int32{11, 12, 13, 14, 10}
	downSymbols := []int32{11, 12, 13, 14, 10}

	return &proto.Response{
		SpinResponse: &proto.SpinResponse{
			Res: &proto.BaseResult{
				Win:      float32(totalWin),
				TotalWin: float32(totalWin),
				Up:       upSymbols,
				Reel:     api.GridConvert(grid),
				Dn:       downSymbols,
				LineWins: api.LinePayConvert(&lineWins),
			},
			Bonus: &proto.SpinResponse_Foodie{Foodie: &proto.FoodieBonus{Free: &proto.FreeSpin{
				Win:    1,
				Remain: 2,
				Max:    3,
			}}},
		},
		State: "thisistempstate",
	}, nil
}
