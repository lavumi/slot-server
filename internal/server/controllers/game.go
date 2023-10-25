package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"slot-server/internal/server/forms"
	"slot-server/internal/server/models"
	"slot-server/internal/slot"
)

type Game struct {
	Slot *slot.Client
}

func (g *Game) Spin(c *gin.Context) {

	var req forms.SpinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//c.JSON(400, gin.H{"msg": err})
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		//SendError(c, 400, err.Error())
		return
	}

	user := c.MustGet("user").(models.User)
	//key := session.Get("key").(string)

	log.Printf("UserID : %s | Cash : %f", user.UUID, user.Cash)
	user.Cash += 1000

	spin, additionalInfo, err := g.Slot.RequestSpin(0, req.BetCash, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	spinObject := make(map[string]interface{})
	err = json.Unmarshal(spin, &spinObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userAfter := c.MustGet("user").(models.User)
	log.Printf("UserID : %s | Cash : %f", userAfter.UUID, userAfter.Cash)
	cashAfter := user.Cash + additionalInfo.CashDiff

	c.JSON(http.StatusOK, forms.SpinResponse{
		SpinResult: spinObject,
		After:      cashAfter,
		Before:     user.Cash,
	})

}
