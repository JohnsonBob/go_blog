package excel

import "go_blog/pkg/setting"

const EXT = ".xlsx"

func GetExcelFullUrl(name string) string {
	return setting.Config.App.PrefixUrl + "/" + GetExcelPath() + name
}
func GetExcelPath() string {
	return setting.Config.App.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.Config.App.RuntimeRootPath + GetExcelPath()
}
