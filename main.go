package main

import (
	"fmt"
	"go_blog/pkg/setting"
	"go_blog/routers"
	"log"
	"net/http"
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
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
