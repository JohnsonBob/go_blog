package main

import (
	"context"
	"fmt"
	"go_blog/models"
	"go_blog/pkg"
	"go_blog/pkg/gredis"
	"go_blog/pkg/logging"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"go_blog/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	setting.Setup()
	gredis.SetUp()
	models.Setup()
	logging.Setup()

	engine := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Config.Server.HttpPort),
		Handler:        engine,
		ReadTimeout:    setting.Config.Server.ReadTimeout,
		WriteTimeout:   setting.Config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			util.Printf("Listen: %s\n", err)
		}
	}()

	//启动定时器
	go pkg.StartClearDataBaseCron()

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
