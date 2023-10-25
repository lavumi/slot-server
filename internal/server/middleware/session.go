package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"slot-server/internal/db"
	"slot-server/internal/server/models"
)

func SessionHandler(rds *db.PseudoRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey := c.GetHeader("session-key")
		sessionString, err := rds.Get(sessionKey).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}

		session := models.Session{}

		//sessionData := make(map[string]string)
		err = json.Unmarshal([]byte(sessionString), &session)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.Set("user", session.User)
		log.Printf("session Check %v", session.User)
		c.Next()
	}
}
