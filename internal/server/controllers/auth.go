package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"slot-server/internal/db"
	"slot-server/internal/server/configs"
	"slot-server/internal/server/forms"
	"slot-server/internal/server/models"
	"time"
)

type Auth struct {
	Redis *db.PseudoRedis
}

func (a *Auth) Guest(c *gin.Context) {

	initCash := configs.InitCash

	userInfo := models.User{
		UUID: uuid.NewString(),
		Cash: initCash,
	}

	sessionKey := uuid.NewString()

	session := models.Session{
		User:       userInfo,
		Key:        sessionKey,
		UpdateTime: time.Now().String(),
	}

	if marshal, err := json.Marshal(session); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		a.Redis.Set(sessionKey, string(marshal), 3600*time.Second)
	}

	res := forms.ResGuest{
		Key:  sessionKey,
		Id:   userInfo.UUID,
		Cash: initCash,
	}

	c.JSON(http.StatusOK, res)
}

//func (a *Auth) upsertSession(key uuid.UUID, user models.User) error {
//	session := models.Session{
//		User:       user,
//		Key:        key.String(),
//		UpdateTime: time.Now().String(),
//	}
//
//	sessionStr, err := json.Marshal(session)
//	if err != nil {
//		return err
//	}
//
//	a.Redis.Set(key.String(), string(sessionStr), time.Hour)
//	return nil
//}
