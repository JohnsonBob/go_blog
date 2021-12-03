package models

import (
	"time"
)

type Article struct {
	Model
	TagId         int    `json:"tag_id"`
	Tag           *Tag   `json:"tag"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"create_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url" gorm:"type:char(255);comment:封面图片地址"`
}

func GetArticle(id int) (article Article) {
	db.Preload("Tag").Where("id = ?", id).First(&article)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func ExistArticleByName(name string) bool {
	var article Article
	db.Select("id").Where("name = ?", name).First(&article)
	return article.ID > 0
}

func ExistArticleById(id int) bool {
	var count int64 = 0
	db.Model(&Article{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func AddArticle(article *Article) {
	article.CreatedOn = time.Now().Unix()
	db.Create(article)
}

func EditArticle(id int, article *Article) {
	article.ID = id
	article.ModifiedOn = time.Now().Unix()
	db.Model(article).Updates(*article)
}

func DeleteArticle(id int) {
	db.Model(&Article{}).Delete("id = ?", id)
}

func CleanAllArticle() {
	db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Article{})
}
