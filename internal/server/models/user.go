package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UUID string  `json:"uuid"`
	Cash float64 `json:"cash"`
}

type UserModel struct {
	db *mongo.Database
}

func InitUserModel(db *mongo.Database) *UserModel {
	return &UserModel{db: db}
}
