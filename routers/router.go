package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/pkg/middleware"
	"go_blog/pkg/setting"
	"go_blog/routers/api"
	v1 "go_blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	engine.POST("/auth", api.GetAuth)

	group := engine.Group("/api/v1")
	group.Use(middleware.JWT())
	{
		//获取标签列表
		group.GET("/tags", v1.GetTags)
		//新建标签
		group.POST("/tags", v1.AddTag)
		//更新指定标签
		group.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		group.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		group.GET("/articles", v1.GetArticles)
		//获取指定文章
		group.GET("/articles/:id", v1.GetArticle)
		//新建文章
		group.POST("/articles", v1.AddArticle)
		//更新指定文章
		group.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		group.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return engine
}
