package excel

import "go_blog/pkg/setting"

const EXT = ".xlsx"
const Sheet = "标签信息"

func GetExcelFullUrl(name string) string {
	return setting.Config.App.PrefixUrl + "/" + GetExcelPath() + name
}
func GetExcelPath() string {
	return setting.Config.App.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.Config.App.RuntimeRootPath + GetExcelPath()
}
