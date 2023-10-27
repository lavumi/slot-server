package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"slot-server/internal/db"
	"slot-server/internal/server/models"
	"time"
)

func SessionHandler(rds *db.PseudoRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey := c.GetHeader("session-key")
		if sessionKey == "" {
			c.Next()
			return
		}
		sessionString, err := rds.Get(sessionKey).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}

		session := models.Session{}
		err = json.Unmarshal([]byte(sessionString), &session)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.Set("user", session.User)
		//log.Printf("session Check %v", session.User)
		c.Next()
	}
}

func SaveSession(rds *db.PseudoRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey := c.GetHeader("session-key")
		if sessionKey == "" {
			c.Next()
			return
		}

		user := c.MustGet("user").(models.User)
		if marshal, err := json.Marshal(user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "UserDataError|" + err.Error(),
			})
			return
		} else {
			rds.Set(sessionKey, string(marshal), 600*time.Second)
		}
		c.Next()
	}
}
