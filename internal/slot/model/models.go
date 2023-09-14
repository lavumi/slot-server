package model

type SpinInput struct {
	Id        int         `json:"id"`
	BetCash   float64     `json:"bet"`
	BetLine   int         `json:"line"`
	PrevState interface{} `json:"prevState,omitempty"`
}

type SpinOutput struct {
	Win         float64      `json:"win"`
	TotalWin    float64      `json:"tw"`
	Symbols     [][]int      `json:"s"`
	UpSymbols   []int        `json:"us"`
	DownSymbols []int        `json:"ds"`
	LineWins    []AllLineWin `json:"lp"`
	BonusWins   []BonusWin   `json:"bn,omitempty"`
	NextProcess interface{}  `json:"next"`
}

type LineWin struct {
	Win      float64 `json:"win"`
	Position []int   `json:"pos"`
}

type AllLineWin struct {
	Win      float64 `json:"win"`
	Position [][]int `json:"pos"`
}

type BonusWin interface{}
