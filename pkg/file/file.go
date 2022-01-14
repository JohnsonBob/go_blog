package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(filePath string) (int64, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	} else {
		return stat.Size(), err
	}
}

func GetSizeFromFile(f multipart.File) (int64, error) {
	content, err := ioutil.ReadAll(f)
	return int64(len(content)), err
}

func GetExt(filePath string) string {
	return path.Ext(filePath)
}

func CheckNotExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return os.IsNotExist(err)
}

func CheckPermission(filePath string) bool {
	_, err := os.Stat(filePath)
	return os.IsPermission(err)
}

func MkDir(filePath string) error {
	return os.MkdirAll(filePath, os.ModePerm)
}

func IsNotExistMkDir(filePath string) error {
	exist := CheckNotExist(filePath)
	if exist {
		err := MkDir(filePath)
		return err
	}
	return nil
}

func Open(filePath string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(filePath, flag, perm)
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	if CheckPermission(src) {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	if IsNotExistMkDir(src) != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
