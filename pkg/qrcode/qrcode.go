package qrcode

import (
	"github.com/boombuler/barcode/qr"
	"go_blog/pkg/file"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	ExtJpg = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		url, width, height, ExtJpg, level, mode,
	}
}

func GetQrCodePath() string {
	return setting.Config.App.QrCodeSavePath
}

func GetQrCodeFullPath() string {
	return setting.Config.App.RuntimeRootPath + setting.Config.App.QrCodeSavePath
}

func GetQrCodeFullUrl(name string) string {
	return setting.Config.App.PrefixUrl + "/" + setting.Config.App.QrCodeSavePath + "/" + name
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}
func (q *QrCode) CheckEncodeExist(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	return file.CheckNotExist(src)
}

/*func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	encode, err := qr.Encode(q.URL, q.Level, q.Mode)
	if err != nil {
		return "", "", err
	}

	scale, err := barcode.Scale(encode, q.Width, q.Height)
}
*/
