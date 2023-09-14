package foodie

import (
	"slot-server/internal/slot/module"
)

type state struct {
	module.State
	RemainFreeSpin int
	TotalFreeSpin  int
	FreeSpinWin    float64
}

type Bonus struct {
	FreeSpin string `json:"freeSpin,omitempty"`
}
