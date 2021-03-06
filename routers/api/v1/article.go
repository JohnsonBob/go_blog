package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/app"
	"go_blog/models"
	"go_blog/pkg/e"
	"go_blog/pkg/qrcode"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"go_blog/service/article_service"
	"go_blog/service/cache_service"
)

// GetArticle 获取单个文章
func GetArticle(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	id := context.Param("id")

	valid := validation.Validation{}
	valid.Numeric(id, "id").Message("ID必须为数字")

	code := e.InvalidParams
	var data interface{}
	if valid.HasErrors() {
		util.PrintLog(&valid)
		response.ResponseWithMessage(code, valid.Errors[0].Message, data)
		return
	}
	idInt := com.StrTo(id).MustInt()
	if models.ExistArticleById(idInt) {
		dataTemp, err := article_service.GetOne(idInt)
		if err != nil {
			code = e.ErrorGetArticleFail
			response.Response500(code, data)
			return
		}
		code = e.SUCCESS
		data = dataTemp
	} else {
		code = e.ErrorNotExistArticle
	}
	response.Response(code, data)

}

// GetArticles 获取多个文章
func GetArticles(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	data := make(map[string]interface{})
	title := context.Query("title")
	tagId := context.Query("tag_id")
	state := context.Query("state")

	valid := validation.Validation{}
	valid.Numeric(tagId, "tag_id").Message("tag_id必须为数字")
	valid.Numeric(state, "state").Message("state必须为数字")
	valid.Range(com.StrTo(state).MustInt(), 0, 1, "state").Message("state只允许0或1")

	if valid.HasErrors() {
		util.PrintLog(&valid)
		response.ResponseWithMessage(e.InvalidParams, valid.Errors[0].Message, nil)
		return
	}
	article := cache_service.Article{}
	if title != "" {
		article.Title = &title
	}
	if tagId != "" {
		iTagId := com.StrTo(tagId).MustInt()
		article.TagId = &iTagId

	}
	if state != "" {
		iState := com.StrTo(state).MustInt()
		article.State = &iState
	}
	page := util.GetPage(context)
	article.PageNum = &page
	article.PageSize = &setting.Config.App.PageSize
	all, err := article_service.GetAll(&article)
	if err != nil {
		response.Response500(e.ErrorGetArticleFail, data)
		return
	}
	data["lists"] = all
	data["total"] = len(*all)
	response.Response(e.SUCCESS, data)
}

//AddArticle 新增文章
func AddArticle(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	article := models.Article{}
	err := context.Bind(&article)
	code := e.InvalidParams
	if err != nil {
		util.Printf(err.Error())
		response.Response(code, nil)
		return
	}

	valid := validation.Validation{}
	valid.Required(article.Title, "title").Message("标题不能为空")
	valid.Required(article.Desc, "desc").Message("简述不能为空")
	valid.Required(article.Content, "content").Message("内容不能为空")
	valid.Required(article.CreatedBy, "created_by").Message("创建人不能为空")
	valid.Required(article.CoverImageUrl, "cover_image_url").Message("封面图片地址不能为空")
	valid.MaxSize(article.CoverImageUrl, 255, "cover_image_url").Message("封面图片地址最长为255字符")
	valid.Range(article.State, 0, 1, "state").Message("状态只允许0或1")

	if !valid.HasErrors() {
		if !models.ExistTagById(article.TagId) {
			code = e.ErrorNotExistTag
		} else {
			models.AddArticle(&article)
			code = e.SUCCESS
		}
	} else {
		util.PrintLog(&valid)
		response.ResponseWithMessage(code, valid.Errors[0].Message, nil)
		return
	}
	response.Response(code, nil)
}

// EditArticle 修改文章
func EditArticle(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}

	article := models.Article{}
	code := e.InvalidParams
	id := context.Param("id")

	err := context.Bind(&article)
	if err != nil {
		util.Printf(err.Error())
		response.Response(code, nil)
		return
	}

	valid := validation.Validation{}
	valid.Numeric(id, "id").Message("id必须为数字")
	valid.Required(article.Title, "title").Message("标题不能为空")
	valid.MaxSize(article.Title, 100, "title").Message("标题最长为100字符")
	valid.Required(article.Desc, "desc").Message("简述不能为空")
	valid.MaxSize(article.Desc, 255, "desc").Message("简述最长为255字符")
	valid.Required(article.Content, "content").Message("内容不能为空")
	valid.MaxSize(article.Content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(article.ModifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(article.ModifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.Required(article.CreatedBy, "created_by").Message("创建人不能为空")
	valid.Range(article.State, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(article.CoverImageUrl, "cover_image_url").Message("封面图片地址不能为空")
	valid.MaxSize(article.CoverImageUrl, 255, "cover_image_url").Message("封面图片地址最长为255字符")

	if !valid.HasErrors() {
		id := com.StrTo(id).MustInt()
		if !models.ExistArticleById(id) {
			code = e.ErrorNotExistArticle
		} else {
			if !models.ExistTagById(article.TagId) {
				code = e.ErrorNotExistTag
			} else {
				models.EditArticle(id, &article)
				code = e.SUCCESS
			}
		}
		response.Response(code, nil)
	} else {
		util.PrintLog(&valid)
		response.ResponseWithMessage(code, valid.Errors[0].Message, nil)
	}

}

// DeleteArticle 删除文章
func DeleteArticle(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	id := context.Param("id")
	valid := validation.Validation{}
	valid.Numeric(id, "id").Message("id必须为数字")
	code := e.InvalidParams

	if !valid.HasErrors() {
		id := com.StrTo(id).MustInt()
		if models.ExistArticleById(id) {
			code = e.SUCCESS
			models.DeleteArticle(id)
		} else {
			code = e.ErrorNotExistArticle
		}
		response.Response(code, nil)
	} else {
		response.ResponseWithMessage(code, valid.Errors[0].Message, nil)
	}
}

const (
	QRCODE_URL = "https://www.lingdian.site/"
)

func GenerateArticlePoster(context *gin.Context) {
	response := app.BaseResponse{Ctx: context}
	qrCode := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)

	posterName := article_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qrCode.URL) + qrCode.GetQrCodeExt()
	articlePoster := article_service.NewArticlePoster(posterName, &models.Article{}, *qrCode)

	articlePosterBgService := article_service.ArticlePosterBg{BgName: "bg.jpg", ArticlePoster: articlePoster,
		Rect: &article_service.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		Pt: &article_service.Pt{
			X: 125,
			Y: 298,
		}}

	name, _, err := articlePosterBgService.Generate()

	if err != nil {
		response.Response(e.ERROR, nil)
		return
	}
	data := make(map[string]interface{})
	data["qr_url"] = qrcode.GetQrCodeFullUrl(name)
	data["qr_save_url"] = qrcode.GetQrCodePath() + name
	response.Response(e.SUCCESS, data)

}
