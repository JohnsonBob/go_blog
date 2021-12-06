package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/pkg/e"
	"go_blog/pkg/upload"
	"go_blog/pkg/util"
	"net/http"
)

func UploadImage(context *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := context.Request.FormFile("image")
	if err != nil {
		util.Println(err)
		code = e.ERROR
		context.JSON(http.StatusOK, e.GetDefault(code, e.GetMsg(code), data))
		return
	}

	if image == nil {
		code = e.InvalidParams
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSizeFromFile(file) {
			code = e.ErrorUploadCheckImageFormat
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				util.Println(err)
				code = e.ErrorUploadCheckImageFail
			} else if err := context.SaveUploadedFile(image, src); err != nil {
				util.Println(err)
				code = e.ErrorUploadSaveImageFail
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	context.JSON(http.StatusOK, e.GetDefault(code, e.GetMsg(code), data))
}
