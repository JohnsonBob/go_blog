package cache_service

import (
	"go_blog/pkg/e"
	"strconv"
	"strings"
)

type Article struct {
	ID       int
	TagId    int
	State    int
	PageNum  int
	PageSize int
}

func GetArticleKey(article *Article) string {
	return e.CacheArticle + "_" + strconv.Itoa(article.ID)
}

func GetArticlesKey(article *Article) string {
	keys := []string{
		e.CacheArticle,
		"LIST",
	}
	if article.ID > 0 {
		_ = append(keys, strconv.Itoa(article.ID))
	}
	if article.TagId > 0 {
		_ = append(keys, strconv.Itoa(article.TagId))
	}
	if article.State > 0 {
		_ = append(keys, strconv.Itoa(article.State))
	}
	if article.PageNum > 0 {
		_ = append(keys, strconv.Itoa(article.PageNum))
	}
	if article.PageSize > 0 {
		_ = append(keys, strconv.Itoa(article.PageSize))
	}
	return strings.Join(keys, "_")
}
