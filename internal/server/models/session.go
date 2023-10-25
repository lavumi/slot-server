package models

type Session struct {
	User
	Key        string `bson:"key"`
	UpdateTime string `bson:"update-time"`
}
