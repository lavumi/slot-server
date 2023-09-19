package foodie

import (
	"slot-server/internal/slot/model"
	"slot-server/internal/slot/module"
)

type ParSheet struct {
	Name     string         `toml:"name"`
	Id       int            `toml:"id"`
	Bets     []int          `toml:"bets"`
	Lines    [][]int        `toml:"lines"`
	Column   int            `toml:"column"`
	Strips   []module.Strip `toml:"strip"`
	Pays     []model.Line   `toml:"line"`
	FreeSpin model.Scatter  `toml:"freeSpin"`
}
