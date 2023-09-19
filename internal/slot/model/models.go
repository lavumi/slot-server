package model

type Bonus int32

const (
	FreeSpin Bonus = iota
	Jackpot
	Pot
)

type Line struct {
	Symbol int       `toml:"symbol,omitempty"`
	Payout []float32 `toml:"payout,omitempty"`
}

type Scatter struct {
	Symbol    int       `toml:"symbol"`
	Payout    []float32 `toml:"payout,omitempty"`
	BonusType Bonus     `toml:"type,omitempty"`
	Bonus     []int     `toml:"bonus,omitempty"`
}
