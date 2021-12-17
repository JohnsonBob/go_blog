package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/pkg/excel"
	"go_blog/pkg/middleware"
	"go_blog/pkg/setting"
	"go_blog/pkg/upload"
	"go_blog/routers/api"
	v1 "go_blog/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.Config.Server.RunMode)

	engine.POST("/auth", api.GetAuth)
	engine.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	engine.StaticFS("/export", http.Dir(excel.GetExcelFullPath()))

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
		//导出tag
		group.GET("/tags/export", v1.ExportTag)
		//导入tag
		group.POST("/tags/import", v1.ImportTag)

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
		//上传图片
		group.POST("/upload", v1.UploadImage)
	}

	return engine
}
