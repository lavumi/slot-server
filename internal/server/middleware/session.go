package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"slot-server/internal/server/models"
)

func SessionHandler(sessionModel models.SessionModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey := c.GetHeader("session-Key")
		sessionUUID, err := uuid.Parse(sessionKey)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sessionData, err := sessionModel.GetSession(sessionUUID)
		if err != nil {
			log.Printf("get session error : %s", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("uid", sessionData.UUID)
		c.Set("cash", sessionData.Cash)

		log.Printf("session Check %v, %f", sessionData.User, sessionData.Cash)
		c.Next()
	}
}

//func SessionHandler
