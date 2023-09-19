package foodie

import (
	"slot-server/internal/slot/module"
)

type state struct {
	module.State
	RemainFreeSpin uint32
	TotalFreeSpin  uint32
	FreeSpinWin    float32
}
