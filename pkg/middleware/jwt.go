package middleware

import (
	"github.com/gin-gonic/gin"
	"go_blog/app"
	"go_blog/pkg/e"
	"go_blog/pkg/util"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		response := app.BaseResponse{Ctx: context}
		var code = e.SUCCESS
		authorization := context.GetHeader("Authorization")
		if fields := strings.Fields(authorization); len(fields) < 2 {
			code = e.ErrorAuthCheckTokenFail
		} else {
			token, err := util.ParseToken(fields[1])
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > token.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.SUCCESS {
			response.ResponseWithHttpCode(http.StatusUnauthorized, code, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
