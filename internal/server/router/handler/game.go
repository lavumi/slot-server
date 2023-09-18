package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"slot-server/internal/server/model"
	"slot-server/internal/slot"
)

func Spin(slot *slot.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.SpinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			//c.JSON(400, gin.H{"msg": err})
			fmt.Println(err.Error())
			SendError(c, 400, err.Error())
			return
		}

		spin, state, err := slot.RequestSpin(0, req.BetCash, "")
		if err != nil {
			SendError(c, 400, err.Error())
			return
		}

		log.Printf("spinState : %s\n", state)

		spinObject := make(map[string]interface{})
		err = json.Unmarshal(spin, &spinObject)
		if err != nil {
			SendError(c, 400, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"BaseResponse": model.BaseResponse{
				Code:    200,
				Message: "success",
			},
			"SpinOutput": spinObject,
		})
	}
}
