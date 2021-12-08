package cache_service

import (
	"go_blog/pkg/e"
	"strconv"
	"strings"
)

type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

func GetTagKey(tag *Tag) string {
	return e.CacheTag + "_" + strconv.Itoa(tag.ID)
}

func GetTagsKey(tag *Tag) string {
	keys := []string{
		e.CacheTag,
		"LIST",
	}
	if tag.ID > 0 {
		_ = append(keys, strconv.Itoa(tag.ID))
	}
	if tag.Name != "" {
		_ = append(keys, tag.Name)
	}
	if tag.State > 0 {
		_ = append(keys, strconv.Itoa(tag.State))
	}
	if tag.PageNum > 0 {
		_ = append(keys, strconv.Itoa(tag.PageNum))
	}
	if tag.PageSize > 0 {
		_ = append(keys, strconv.Itoa(tag.PageSize))
	}
	return strings.Join(keys, "_")
}
