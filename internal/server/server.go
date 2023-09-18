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

	slotClient, err := slot.Connect()
	if err != nil {
		panic("login to slot server fail")
	}

	router.InitRouter(r, slotClient)

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
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		database.DisConnect()
		log.Println("timeout of 5 seconds.")
	}
	log.Println("server exiting")
}
