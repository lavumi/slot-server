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

	log.Printf("UserID : %s ", uid)

	spin, state, cash, err := g.Slot.RequestSpin(0, req.BetCash, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("spinState : %s\n", string(state))
	log.Printf("spinState : %f\n", cash)

	spinObject := make(map[string]interface{})
	err = json.Unmarshal(spin, &spinObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, spinObject)

}
