package module

type State struct {
	Bet  float64 `json:"bet,omitempty"`
	Line int     `json:"line,omitempty"`
}

const WildSymbol = 99
