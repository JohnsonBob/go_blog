package app

import (
	"github.com/gin-gonic/gin"
	"go_blog/pkg/e"
	"net/http"
)

type BaseResponse struct {
	Ctx *gin.Context
}

func (response *BaseResponse) Response(errCode int, data interface{}) {
	response.Ctx.JSON(http.StatusOK, e.GetDefault(errCode, e.GetMsg(errCode), data))
}

func (response *BaseResponse) ResponseWithMessage(errCode int, message string, data interface{}) {
	response.Ctx.JSON(http.StatusOK, e.GetDefault(errCode, message, data))
}

func (response *BaseResponse) ResponseWithHttpCode(httpCode int, errCode int, data interface{}) {
	response.Ctx.JSON(httpCode, e.GetDefault(errCode, e.GetMsg(errCode), data))
}
