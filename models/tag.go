package models

import (
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
	tag.CreatedOn = time.Now().Unix()
	db.Create(tag)
}

func EditTag(id int, tag *Tag) {
	tag.ModifiedOn = time.Now().Unix()
	tag.ID = id
	db.Model(tag).Updates(*tag)
}

func DeleteTag(id int) {
	db.Model(&Tag{}).Delete("id = ?", id)
}

func CleanAllTag() {
	db.Unscoped().Where("deleted_at is not null").Delete(&Tag{})
}
