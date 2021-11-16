package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/pkg/setting"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	engine.GET("/test", func(ctx *gin.Context) {
		ctx.JSONP(200, gin.H{
			"message": "test",
		})
	})
	return engine
}
