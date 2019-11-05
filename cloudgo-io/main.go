package main

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	engine := gin.Default()
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(static.Serve("/", static.LocalFile("./static", true)))
	engine.Use(EncodingHandler())
	engine.GET("/user-list", JsonHandler)
	engine.GET("/unknown", UnknownHandler)
	engine.POST("/form", FormHandler)

	server := &http.Server{
		Addr: "localhost:2345",
		Handler: engine,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	select {
	case sig := <-sc:
		atomic.StoreInt32(&state, 0)
		fmt.Printf("Signal %s", sig.String())
	}

	if err := engine.Run(); err != nil {
		panic(err)
	}

	os.Exit(int(atomic.LoadInt32(&state)))
}