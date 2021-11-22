package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_blog/models"
	"go_blog/pkg/e"
	"go_blog/pkg/util"
	"net/http"
)

type user struct {
	UserName string `valid:"Required; MaxSize(50)" json:"user_name"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetAuth(ctx *gin.Context) {
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
		ctx.JSON(http.StatusOK, e.GetDefault(code, e.GetMsg(code), data))
	} else {
		util.PrintLog(&valid)
		ctx.JSON(http.StatusOK, e.GetDefault(code, valid.Errors[0].Message, nil))
	}
}
