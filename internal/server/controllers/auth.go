package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"slot-server/internal/server/configs"
	"slot-server/internal/server/forms"
	"slot-server/internal/server/models"
)

type Auth struct {
}

func (a *Auth) Guest(c *gin.Context) {
	initCash := configs.InitCash

	userInfo := models.User{
		UUID: uuid.NewString(),
		Cash: initCash,
	}

	sessionKey := uuid.New()

	err := models.UpsertSession(sessionKey, userInfo)
	if err != nil {
		return
	}

	res := forms.ResGuest{
		Key:  sessionKey.String(),
		Id:   userInfo.UUID,
		Cash: initCash,
	}

	c.JSON(http.StatusOK, res)
}
