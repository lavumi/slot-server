package foodie

import (
	"encoding/json"
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

func Spin(req *proto.Request) (*proto.SpinResponse, *model.Error) {
	s := state{}
	if req.PrevState != nil {
		if err := json.Unmarshal(req.PrevState, &s); err != nil {
			return nil, &model.Error{
				Code:    model.ERR_UNKNOWN,
				Message: "FailToParseStateData",
			}
		}
	}

	freeSpinMode := s.TotalFreeSpin > 0

	grid := module.GenRandomGrid(ps.Strips, ps.Column)
	upSymbols := []int32{11, 12, 13, 14, 10}
	downSymbols := []int32{11, 12, 13, 14, 10}

	lineWins := module.AllLineWins(grid, ps.Pays, req.BetCash)
	freeSpinWin := module.ScatterWin(grid, ps.FreeSpin, req.BetCash*50)

	var winCash float32 = 0
	for _, win := range lineWins {
		winCash += win.Win
	}

	var totalWin float32 = 0
	if freeSpinMode == false {
		totalWin = winCash
	} else {
		s.FreeSpinWin += winCash
		s.RemainFreeSpin--
	}

	if freeSpinWin != nil {
		s.RemainFreeSpin += uint32(freeSpinWin.BonusParam)
		s.TotalFreeSpin += uint32(freeSpinWin.BonusParam)
	}

	currentState, err := json.Marshal(s)
	if err != nil {
		return nil, &model.Error{
			Code:    model.ERR_UNKNOWN,
			Message: "FailToParseStateData",
		}
	}

	return &proto.SpinResponse{
		Result: &proto.SpinResult{
			Res: &proto.BaseResult{
				Win:      winCash,
				TotalWin: totalWin,
				Up:       upSymbols,
				Reel:     api.GridConvert(grid),
				Dn:       downSymbols,
				LineWins: lineWins,
			},
			Bonus: &proto.SpinResult_Foodie{Foodie: &proto.FoodieBonus{
				Free: &proto.FreeSpin{
					Win:    s.FreeSpinWin,
					Remain: s.RemainFreeSpin,
					Max:    s.TotalFreeSpin,
				},
			}},
		},
		State: currentState,
		Cash:  winCash - req.BetCash*50,
	}, nil
}
