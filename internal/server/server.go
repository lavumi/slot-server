package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"slot-server/internal/database"
	"slot-server/internal/server/router"
	"slot-server/internal/slot"
	"syscall"
	"time"
)

var r *gin.Engine
var srv *http.Server

func Run() {

	initialize()

	//여기를 어떻게 예쁘게 할지 고민이 좀 필요합니다.
	m := slot.Initialize()
	router.InitRouter(r, m)

	run()
}

func initialize() {
	//gin.SetMode(gin.ReleaseMode)
	r = gin.New()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Panic("set trusted proxies fail")
		return
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./web", false)))
}

func run() {
	srv = &http.Server{
		Addr:        ":8081",
		Handler:     r,
		ReadTimeout: 10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
		return
	}
	// Wait for interrupt signal to gracefully shut down the srv with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need added it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		database.DisConnect()
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
