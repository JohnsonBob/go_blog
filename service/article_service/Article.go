package article_service

import (
	"encoding/json"
	"go_blog/models"
	"go_blog/pkg/gredis"
	"go_blog/service/cache_service"
)

func GetOne(id int) (article *models.Article, err error) {
	cacheArticle := cache_service.Article{ID: id}
	key := cacheArticle.GetArticleKey()
	if gredis.IsExists(key) {
		var data []byte
		data, err = gredis.Get(key)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &article)
		return
	}

	article = models.GetArticle(id)
	err = gredis.Set(key, article, 1200)
	return
}

//func ExistArticleById(id int) (bool, err error) {
//	cacheArticle := cache_service.Article{ID: id}
//	key := cacheArticle.GetArticleKey()
//
//}
