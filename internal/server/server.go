package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"slot-server/internal/server/controllers"
	"slot-server/internal/server/middleware"
	"slot-server/internal/slot"
	"syscall"
	"time"
)

var r *gin.Engine
var srv *http.Server

func initGin() {
	r = gin.New()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic("set trusted proxies fail")
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./web", false)))

	slotClient, err := slot.Connect()
	if err != nil {
		panic("login to gameController server fail")
	}

	authController := controllers.Auth{}
	gameController := controllers.Game{
		Slot: slotClient,
	}

	apiRouter := r.Group("/api")
	{
		authRouter := apiRouter.Group("/auth")
		{
			authRouter.POST("/guest", authController.Guest)
		}
		gameRouter := apiRouter.Group("/game")
		gameRouter.Use(middleware.SessionHandler)
		{

			gameRouter.POST("/:id/enter")
			gameRouter.POST("/:id/spin", gameController.Spin)
			gameRouter.POST("/:id/collect")
			gameRouter.GET("/:id/info")
		}
	}
}

func Run() {

	err := godotenv.Load(".web.dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initGin()

	srv = &http.Server{
		Addr:        ":8081",
		Handler:     r,
		ReadTimeout: 10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
		return
	}

	//region [ Gracefully Shutdown ]
	quit := make(chan os.Signal)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need added it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		//db.DisConnect()
		log.Println("timeout of 5 seconds.")
	}
	log.Println("server exiting")
	//endregion
}
