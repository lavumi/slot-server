package models

type User struct {
	UUID string  `bson:"uuid"`
	Cash float64 `bson:"cash"`
}
