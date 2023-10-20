package foodie

type state struct {
	RemainFreeSpin uint32  `json:"rf,omitempty"`
	TotalFreeSpin  uint32  `json:"tf,omitempty"`
	FreeSpinWin    float32 `json:"fw,omitempty"`
}
