package slot00

import (
	"slot-server/internal/slot/module"
)

type freeSpin struct {
	Symbol  int      `toml:"symbol,omitempty"`
	Feature []string `toml:"feature,omitempty"`
}

type ParSheet struct {
	Name     string         `toml:"name"`
	Id       int            `toml:"id"`
	Bets     []int          `toml:"bets"`
	Lines    [][]int        `toml:"lines"`
	Column   int            `toml:"column"`
	Strips   []module.Strip `toml:"strip"`
	Pays     []module.Pay   `toml:"pays"`
	FreeSpin freeSpin       `toml:"freeSpin"`
}
