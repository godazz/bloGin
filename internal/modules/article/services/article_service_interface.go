package services

import (
	articleResponse "github.com/godazz/bloGin/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() articleResponse.Articles
	GetStoriesArticles() articleResponse.Articles
}
