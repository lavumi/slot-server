package models

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	UUID string  `json:"uuid"`
	Cash float32 `json:"cash"`
}

type UserModel struct {
	col *mongo.Collection
}

func (m *UserModel) Initialize(db *mongo.Database) {
	m.col = db.Collection("User")
}
