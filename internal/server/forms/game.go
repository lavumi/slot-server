package forms

type SpinRequest struct {
	BaseRequest
	SpinInput
}

type SpinInput struct {
	Id        uint32  `json:"id"`
	BetCash   float32 `json:"bet"`
	PrevState string  `json:"prevState,omitempty"`
}
