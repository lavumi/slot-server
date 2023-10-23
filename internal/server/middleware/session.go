package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"slot-server/internal/server/models"
)

func SessionHandler(c *gin.Context) {
	sessionKey := c.GetHeader("Session-Key")
	sessionUUID, err := uuid.Parse(sessionKey)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	sessionData, err := models.GetSession(sessionUUID)
	if err != nil {
		log.Printf("get session error : %s", err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("uid", sessionData.UUID)
	c.Set("cash", sessionData.Cash)

	log.Printf("Session Check %v, %f", sessionData.User, sessionData.Cash)
	c.Next()
}
