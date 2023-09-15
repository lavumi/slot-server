package model

type SpinRequest struct {
	BaseRequest
	SpinInput
}

type SpinResponse struct {
	BaseResponse
	SpinOutput
}

type SpinInput struct {
	Id        uint32  `json:"id"`
	BetCash   float32 `json:"bet"`
	BetLine   int     `json:"line"`
	PrevState string  `json:"prevState,omitempty"`
}

type AllLineWin struct {
	Win      float64 `json:"win"`
	Position [][]int `json:"pos"`
}

type SpinOutput struct {
	Win         float64           `json:"win"`
	TotalWin    float64           `json:"tw"`
	Symbols     [][]int           `json:"s"`
	UpSymbols   []int             `json:"us"`
	DownSymbols []int             `json:"ds"`
	LineWins    []AllLineWin      `json:"lp"`
	BonusWins   map[string]string `json:"bn,omitempty"`
	NextProcess interface{}       `json:"next"`
}
