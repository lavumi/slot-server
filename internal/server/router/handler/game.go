package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"slot-server/internal/server/model"
	"slot-server/internal/slot"
)

func Spin(m *slot.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.SpinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			//c.JSON(400, gin.H{"msg": err})
			fmt.Println(err.Error())
			SendError(c, 400, err.Error())
			return
		}

		//todo
		//spinOutput, err := m.Spin(req.Id, req.BetCash, req.PrevState)
		//if err != nil {
		//	SendError(c, 400, err.Error())
		//	return
		//}

		c.JSON(200, gin.H{
			"BaseResponse": model.BaseResponse{
				Code:    200,
				Message: "success",
			},
			"SpinOutput": nil,
		})
	}
}
