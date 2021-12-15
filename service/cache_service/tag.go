package cache_service

import (
	"go_blog/pkg/e"
	"strconv"
	"strings"
)

type Tag struct {
	ID    *int
	Name  *string
	State *int

	PageNum  *int
	PageSize *int
}

func (tag *Tag) GetTagKey() string {
	return e.CacheTag + "_" + strconv.Itoa(*tag.ID)
}

func (tag *Tag) GetTagsKey() string {
	keys := []string{
		e.CacheTag,
		"LIST",
	}
	if tag.ID != nil {
		keys = append(keys, strconv.Itoa(*tag.ID))
	}
	if tag.Name != nil {
		keys = append(keys, *tag.Name)
	}
	if tag.State != nil {
		keys = append(keys, strconv.Itoa(*tag.State))
	}
	if tag.PageNum != nil {
		keys = append(keys, strconv.Itoa(*tag.PageNum))
	}
	if tag.PageSize != nil {
		keys = append(keys, strconv.Itoa(*tag.PageSize))
	}
	return strings.Join(keys, "_")
}
