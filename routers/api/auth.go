package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_blog/app"
	"go_blog/models"
	"go_blog/pkg/e"
	"go_blog/pkg/util"
)

type user struct {
	UserName string `valid:"Required; MaxSize(50)" json:"user_name"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetAuth(ctx *gin.Context) {
	response := app.BaseResponse{Ctx: ctx}
	user := user{}
	_ = ctx.Bind(&user)
	valid := validation.Validation{}
	ok, _ := valid.Valid(&user)

	code := e.InvalidParams
	data := make(map[string]interface{})
	if ok {
		exists := models.CheckAuth(user.UserName, user.Password)
		if exists {
			token, err := util.GenerateToken(user.UserName, user.Password)
			if err == nil {
				code = e.SUCCESS
				data["token"] = token
			} else {
				code = e.ErrorAuthToken
			}
		} else {
			code = e.ErrorAuth
		}
		response.Response(code, data)
	} else {
		util.PrintLog(&valid)
		response.ResponseWithMessage(code, valid.Errors[0].Message, data)
	}
}
