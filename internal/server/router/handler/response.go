package handler

import (
	"github.com/gin-gonic/gin"
)

//func SendResponse(c *gin.Context, code int, res interface{}) {
//	//
//	var response = model.SpinResponse{
//		BaseResponse: model.BaseResponse{
//			Code:    200,
//			Message: "success",
//		},
//		SpinOutput: res.(model.SpinOutput),
//	}
//	c.JSON(code, response)
//}

func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"msg": message})
}
