package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/app"
	"go_blog/models"
	"go_blog/pkg/e"
	"go_blog/pkg/excel"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"go_blog/service/cache_service"
	"go_blog/service/tag_service"
)

// GetTags 获取多个文章标签
func GetTags(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	data := make(map[string]interface{})
	tag := cache_service.Tag{}
	if name := context.Query("name"); name != "" {
		tag.Name = &name
	}

	var state = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		tag.State = &state
	}
	page := util.GetPage(context)
	tag.PageNum = &page
	tag.PageSize = &setting.Config.App.PageSize

	tags, err := tag_service.GetTags(&tag)
	if err != nil {
		response.Response500(e.ErrorGetArticlesFail, data)
		return
	}
	data["lists"] = tags
	data["total"] = len(*tags)
	response.Response(e.SUCCESS, data)
}

// AddTag 新增文章标签
func AddTag(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
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
		response.Response(code, nil)
	} else {
		response.ResponseWithMessage(code, valid.Errors[0].Message, nil)
	}
}

// EditTag 修改文章标签
func EditTag(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
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
		response.Response(code, nil)
	} else {
		response.ResponseWithMessage(code, valid.Errors[0].Message, nil)

	}

}

// DeleteTag 删除文章标签
func DeleteTag(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
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
		response.Response(code, nil)
	} else {
		response.ResponseWithMessage(code, valid.Errors[0].Message, nil)
	}
}

func ExportTag(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	data := make(map[string]interface{})
	tag := cache_service.Tag{}
	if name := context.Query("name"); name != "" {
		tag.Name = &name
	}

	var state = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		tag.State = &state
	}
	page := util.GetPage(context)
	tag.PageNum = &page
	tag.PageSize = &setting.Config.App.PageSize

	name, err := tag_service.Export(&tag)
	if err != nil {
		response.Response500(e.ErrorExportTagFail, data)
		return
	}
	data["export_url"] = excel.GetExcelFullUrl(name)
	data["export_save_url"] = excel.GetExcelFullPath() + name
	response.Response(e.SUCCESS, data)
}

func ImportTag(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		util.Println(err)
		response.Response(e.InvalidParams, nil)
		return
	}
	err = tag_service.ImportTag(file)
	if err != nil {
		util.Println(err)
		response.Response500(e.ErrorImportTagFail, err)
		return
	}
	response.Response(e.SUCCESS, nil)
}
