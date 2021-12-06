package file

import (
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
