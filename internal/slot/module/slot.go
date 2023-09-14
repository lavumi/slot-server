package module

import "slot-server/internal/slot/model"

type State struct {
	Bet  float64 `json:"bet,omitempty"`
	Line int     `json:"line,omitempty"`
}

type ISlot interface {
	Spin(prev interface{}, bet float64) (*model.SpinOutput, *model.Error)
}

const WildSymbol = 99
