package e

import "github.com/gin-gonic/gin"

func GetDefault(code int, message string, data interface{}) interface{} {
	return gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
}
