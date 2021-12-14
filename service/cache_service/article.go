package cache_service

import (
	"go_blog/pkg/e"
	"strconv"
	"strings"
)

type Article struct {
	ID       *int
	TagId    *int
	State    *int
	PageNum  *int
	PageSize *int
	Title    *string
}

func (article *Article) GetArticleKey() string {
	return e.CacheArticle + "_" + strconv.Itoa(*article.ID)
}

func (article *Article) GetArticlesKey() string {
	keys := []string{
		e.CacheArticle,
		"LIST",
	}
	if article.ID != nil {
		keys = append(keys, strconv.Itoa(*article.ID))
	}
	if article.TagId != nil {
		keys = append(keys, strconv.Itoa(*article.TagId))
	}
	if article.State != nil {
		keys = append(keys, strconv.Itoa(*article.State))
	}
	if article.Title != nil {
		keys = append(keys, *article.Title)
	}
	if article.PageNum != nil {
		keys = append(keys, strconv.Itoa(*article.PageNum))
	}
	if article.PageSize != nil {
		keys = append(keys, strconv.Itoa(*article.PageSize))
	}
	return strings.Join(keys, "_")
}
