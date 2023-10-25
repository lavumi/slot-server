package models

type Session struct {
	User
	Key        string `json:"key"`
	UpdateTime string `json:"update-time"`
}
