package models

type BaseRequest struct {
	Index   int `json:"index"`
	Counter int `json:"counter"`
}
