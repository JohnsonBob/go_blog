package middleware

import (
	"github.com/gin-gonic/gin"
	"go_blog/pkg/e"
	"go_blog/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code = e.SUCCESS
		authorization := context.GetHeader("Authorization")
		if authorization == "" {
			code = e.InvalidParams
		} else {
			token, err := util.ParseToken(authorization)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > token.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.SUCCESS {
			context.JSON(http.StatusUnauthorized, e.GetDefault(code, e.GetMsg(code), nil))
			context.Abort()
			return
		}
		context.Next()
	}
}
