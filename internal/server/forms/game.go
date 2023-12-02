package forms

type EnterRequest struct {
	BaseRequest
	EnterInput
}

type EnterInput struct {
	Id uint32 `json:"id"`
}

type EnterResponse struct {
	GameInfo map[string]interface{} `json:"gameInfo"`
}

type SpinRequest struct {
	BaseRequest
	SpinInput
}

type SpinInput struct {
	Id        uint32  `json:"id"`
	BetCash   float32 `json:"bet"`
	PrevState string  `json:"prevState,omitempty"`
}

type SpinResponse struct {
	SpinResult map[string]interface{} `json:"spin"`
	Before     float64                `json:"before"`
	After      float64                `json:"after"`
}
