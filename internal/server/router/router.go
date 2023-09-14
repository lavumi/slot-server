package router

import (
	"github.com/gin-gonic/gin"
	"slot-server/internal/server/router/handler"
	"slot-server/internal/slot"
	//"slot-crawler/api/router"
)

func InitRouter(r *gin.Engine, manager *slot.Manager) *gin.Engine {
	apiRouter := r.Group("/api")
	{
		//authRouter := apiRouter.Group("/auth")
		//{
		//	authRouter.POST("/login")
		//	authRouter.POST("/logout")
		//}
		gameRouter := apiRouter.Group("/game")
		{
			//gameRouter.POST("/:id/enter")
			gameRouter.POST("/:id/spin", handler.Spin(manager))
			//gameRouter.POST("/:id/collect")
			//gameRouter.GET("/:id/info")
		}
		//testRouter := apiRouter.Group("/test")
		//{
		//	testRouter.POST("/charge")
		//	testRouter.POST("/reset")
		//}
	}
	return r
}
