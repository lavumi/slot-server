package model

type SpinRequest struct {
	BaseRequest
	SpinInput
}

type SpinInput struct {
	Id        uint32  `json:"id"`
	BetCash   float32 `json:"bet"`
	BetLine   int     `json:"line"`
	PrevState string  `json:"prevState,omitempty"`
}
