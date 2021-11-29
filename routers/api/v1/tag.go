package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/models"
	"go_blog/pkg/e"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"net/http"
)

// GetTags 获取多个文章标签
func GetTags(context *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := context.Query("name"); name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(context), setting.Config.App.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTag 新增文章标签
func AddTag(context *gin.Context) {
	tag := models.Tag{}
	_ = context.Bind(&tag)

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Required(tag.CreatedBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(tag.CreatedBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if !models.ExistTagByName(tag.Name) {
			code = e.SUCCESS
			models.AddTag(&tag)
		} else {
			code = e.ErrorExistTag
		}
		context.JSON(http.StatusOK, e.GetDefault(code, e.GetMsg(code), nil))
	} else {
		context.JSON(http.StatusOK, e.GetDefault(code, valid.Errors[0].Message, nil))
	}
}

// EditTag 修改文章标签
func EditTag(context *gin.Context) {
	tag := models.Tag{}
	_ = context.Bind(&tag)
	valid := validation.Validation{}

	id := context.Param("id")
	valid.Numeric(id, "id").Message("id必须为数字")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Required(tag.ModifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(tag.ModifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	code := e.InvalidParams
	if !valid.HasErrors() {
		idInt := com.StrTo(id).MustInt()
		if models.ExistTagById(idInt) {
			code = e.SUCCESS
			models.EditTag(idInt, &tag)
		} else {
			code = e.ErrorNotExistTag
		}
		context.JSON(http.StatusOK, e.GetDefault(code, e.GetMsg(code), nil))
	} else {
		context.JSON(http.StatusOK, e.GetDefault(code, valid.Errors[0].Message, nil))
	}

}

// DeleteTag 删除文章标签
func DeleteTag(context *gin.Context) {
	id := context.Param("id")
	valid := validation.Validation{}
	valid.Numeric(id, "id").Message("id必须为数字")
	code := e.InvalidParams

	if !valid.HasErrors() {
		id := com.StrTo(id).MustInt()
		if models.ExistTagById(id) {
			code = e.SUCCESS
			models.DeleteTag(id)
		} else {
			code = e.ErrorNotExistTag
		}
		context.JSON(http.StatusOK, e.GetDefault(code, e.GetMsg(code), nil))
	} else {
		context.JSON(http.StatusOK, e.GetDefault(code, valid.Errors[0].Message, nil))
	}
}
