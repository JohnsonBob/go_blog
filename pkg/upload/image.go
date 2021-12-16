package upload

import (
	"fmt"
	"go_blog/pkg/file"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"mime/multipart"
	"os"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.Config.App.PrefixUrl + "/" + GetImagePath() + name
}

func GetImagePath() string {
	return setting.Config.App.ImageSavePath
}

func GetImageName(name string) string {
	ext := file.GetExt(name)
	fileName := strings.TrimSuffix(name, ext)
	md5 := util.EncodeMD5(fileName)
	return md5 + ext
}

func GetImageFullPath() string {
	return setting.Config.App.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.Config.App.ImageAllowExts {
		if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
			return true
		}
	}
	return false
}

func CheckImageSize(filePath string) bool {
	size, err := file.GetSize(filePath)
	if err != nil {
		util.Println(err)
		return false
	}
	return size <= setting.Config.App.ImageMaxSize*1024*1024
}

func CheckImageSizeFromFile(f multipart.File) bool {
	size, err := file.GetSizeFromFile(f)
	if err != nil {
		util.Println(err)
		return false
	}

	return size <= setting.Config.App.ImageMaxSize*1024*1024
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
