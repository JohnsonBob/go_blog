package logging

import (
	"fmt"
	"go_blog/pkg/file"
	"go_blog/pkg/setting"
	"log"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.Config.App.RuntimeRootPath, setting.Config.App.LogSavePath)
}

func getLogFileFullPath() string {
	path := getLogFilePath()
	suffixPath := fmt.Sprintf("%s-%s.%s", setting.Config.App.LogSaveName, time.Now().Format(setting.Config.App.TimeFormat), setting.Config.App.LogFileExt)
	return fmt.Sprintf("%s%s", path, suffixPath)
}

func getLogFileName() string {
	suffixPath := fmt.Sprintf("%s-%s.%s", setting.Config.App.LogSaveName, time.Now().Format(setting.Config.App.TimeFormat), setting.Config.App.LogFileExt)
	return suffixPath
}

func openLogFile(filePath string, fileName string) *os.File {
	src := filePath + fileName
	if file.CheckPermission(src) {
		log.Fatalf("file.CheckPermission Permission denied")

	}
	err := file.IsNotExistMkDir(filePath)
	if err != nil {
		log.Fatalf("file.IsNotExistMkDir src: %s, err: %v", filePath, err)
	}

	f, err := file.Open(src, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)

	}
	return f
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
