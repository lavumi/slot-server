package forms

type ReqGuest struct {
}

type ResGuest struct {
	Key  string  `json:"key"`
	Id   string  `json:"id"`
	Cash float64 `json:"cash"`
}
