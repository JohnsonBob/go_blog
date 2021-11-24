package main

import (
	"context"
	"fmt"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"go_blog/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	engine := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        engine,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			util.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	util.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		util.Fatal("Server Shutdown:", err)
	}

	util.Println("Server exiting")
}
