package excel

import "go_blog/pkg/setting"

func GetExcelFullUrl(name string) string {
	return setting.Config.App.ExportSavePath + GetExcelPath() + name
}
func GetExcelPath() string {
	return setting.Config.App.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.Config.App.RuntimeRootPath + GetExcelPath()
}
