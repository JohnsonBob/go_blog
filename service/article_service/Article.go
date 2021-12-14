package article_service

import (
	"encoding/json"
	"go_blog/models"
	"go_blog/pkg/gredis"
	"go_blog/service/cache_service"
)

func GetOne(id int) (article *models.Article, err error) {
	cacheArticle := cache_service.Article{ID: &id}
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

func GetAll(article *cache_service.Article) (articles *[]models.Article, err error) {
	key := article.GetArticlesKey()
	if gredis.IsExists(key) {
		var data []byte
		data, err = gredis.Get(key)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &articles)
		return
	}
	maps := make(map[string]interface{})
	if article.Title != nil {
		maps["title"] = article.Title
	}
	if article.TagId != nil {
		maps["tag_id"] = article.TagId
	}
	if article.State != nil {
		maps["state"] = article.State
	}
	all := models.GetArticles(*article.PageNum, *article.PageSize, maps)
	articles = &all
	err = gredis.Set(key, articles, 1200)
	return
}

//func ExistArticleById(id int) (bool, err error) {
//	cacheArticle := cache_service.Article{ID: id}
//	key := cacheArticle.GetArticleKey()
//
//}
