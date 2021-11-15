package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/pkg/setting"
)

// GetPage /
func GetPage(c *gin.Context) int {
	result := 0
	page, err := com.StrTo(c.Query("page")).Int()
	if err != nil {
		return result
	}
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
