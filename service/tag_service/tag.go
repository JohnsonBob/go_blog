package tag_service

import (
	"encoding/json"
	"go_blog/models"
	"go_blog/pkg/gredis"
	"go_blog/service/cache_service"
)

func GetTags(tag *cache_service.Tag) (tags *[]models.Tag, err error) {
	key := tag.GetTagsKey()
	if gredis.IsExists(key) {
		var data []byte
		data, err = gredis.Get(key)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &tags)
		return
	}
	maps := make(map[string]interface{})
	if tag.Name != nil {
		maps["name"] = tag.Name
	}
	if tag.State != nil {
		maps["state"] = tag.State
	}
	all := models.GetTags(*tag.PageNum, *tag.PageSize, maps)
	tags = &all
	err = gredis.Set(key, tags, 1200)
	return
}
