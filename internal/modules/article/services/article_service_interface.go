package services

import articleModel "github.com/godazz/bloGin/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []articleModel.Article
	GetStoriesArticles() []articleModel.Article
}
