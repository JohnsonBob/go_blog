package article_service

import (
	"go_blog/models"
	"go_blog/pkg/file"
	"go_blog/pkg/qrcode"
	"image"
	"image/draw"
	"image/jpeg"
)

type ArticlePoster struct {
	PosterName string
	*models.Article
	QrCode qrcode.QrCode
}

func NewArticlePoster(posterName string, article *models.Article, qrCode qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{posterName, article, qrCode}

}
func GetPosterFlag() string {
	return "poster"
}

func (articlePoster *ArticlePoster) CheckMergedImage(path string) bool {
	return !file.CheckNotExist(path + articlePoster.PosterName)
}

type ArticlePosterBg struct {
	BgName string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

type Pt struct {
	X int
	Y int
}

func (a *ArticlePosterBg) Generate() (name string, path string, err error) {
	qrFullPath := qrcode.GetQrCodeFullPath()
	qrName, _, err := a.QrCode.Encode(qrFullPath)
	if err != nil {
		return
	}

	if !a.CheckMergedImage(qrFullPath) {
		posterFile, err := file.MustOpen(a.PosterName, qrFullPath)
		bgFile, err := file.MustOpen(a.BgName, qrFullPath)
		qrFile, err := file.MustOpen(qrName, qrFullPath)
		if err != nil {
			return name, path, err
		}
		defer func() {
			_ = posterFile.Close()
			_ = bgFile.Close()
			_ = qrFile.Close()
		}()

		bgImage, err := jpeg.Decode(bgFile)
		qrImage, err := jpeg.Decode(qrFile)
		if err != nil {
			return name, path, err
		}
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
		err = jpeg.Encode(posterFile, jpg, nil)
		if err != nil {
			return name, path, err
		}
	}
	return a.PosterName, qrFullPath, nil
}
