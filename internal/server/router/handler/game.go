package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	model2 "slot-server/internal/server/model"
	"slot-server/internal/slot"
)

func Spin(m *slot.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model2.SpinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			//c.JSON(400, gin.H{"msg": err})
			fmt.Println(err.Error())
			SendError(c, 400, err.Error())
			return
		}
		spinOutput, err := m.Spin(req.SpinInput)
		if err != nil {
			SendError(c, 400, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"BaseResponse": model2.BaseResponse{
				Code:    200,
				Message: "success",
			},
			"SpinOutput": spinOutput,
		})
	}
}
