package models

type UserStateSlot struct {
	UUID  string  `json:"uuid"`
	State []byte  `json:"state"`
	Cash  float64 `json:"cash"`
}
