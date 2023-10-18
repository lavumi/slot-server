package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"slot-server/internal/db"
	"time"
)

type Session struct {
	User
	Key        string `json:"key"`
	UpdateTime string `json:"update-time"`
}

func UpsertSession(key uuid.UUID, user User) error {
	session := Session{
		user,
		key.String(),
		time.Now().String(),
	}

	sessionStr, err := json.Marshal(session)
	if err != nil {
		return err
	}

	db.GetRedis().Set(key.String(), string(sessionStr), time.Hour)
	return nil
}

func GetSession(key uuid.UUID) (*Session, error) {
	if res, err := db.GetRedis().Get(key.String()).Result(); err != nil {
		return nil, err
	} else {
		var s Session
		if err := json.Unmarshal([]byte(res), &s); err != nil {
			return nil, err
		} else {
			return &s, nil
		}
	}

}
