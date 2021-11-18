package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

func ExistTagById(id int) bool {
	var count int64 = 0
	db.Model(&Tag{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func AddTag(tag *Tag) {
	db.Create(tag)
}
func (tag *Tag) BeforeCreate(tx *gorm.DB) error {
	tx.Model(tag).UpdateColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(tx *gorm.DB) error {
	tx.Model(tag).UpdateColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func EditTag(id int, tag *Tag) {
	db.Model(tag).Where("id = ?", id).Updates(*tag)
}
