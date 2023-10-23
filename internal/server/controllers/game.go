package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"slot-server/internal/server/forms"
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

	uid := c.GetString("uid")
	cash := c.GetFloat64("cash")

	log.Printf("UserID : %s | Cash : %f", uid, cash)

	spin, state, diff, err := g.Slot.RequestSpin(0, req.BetCash, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("spinState : %s\n", string(state))

	spinObject := make(map[string]interface{})
	err = json.Unmarshal(spin, &spinObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, forms.SpinResponse{
		SpinResult: spinObject,
		After:      cash + float64(diff),
		Before:     cash,
	})

}
