package models

/*import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model
	TagId      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(tx *gorm.DB) error {
	tx.Model(article).Update("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(tx *gorm.DB) error {
	tx.Model(article).Update("ModifiedOn", time.Now().Unix())

	return nil
}*/
