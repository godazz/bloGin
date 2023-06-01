package services

import (
	articleModel "github.com/godazz/bloGin/internal/modules/article/models"
	articleRepository "github.com/godazz/bloGin/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: articleRepository.New(),
	}
}

func (articleService ArticleService) GetFeaturedArticles() []articleModel.Article {
	return articleService.articleRepository.List(4)
}

func (articleService ArticleService) GetStoriesArticles() []articleModel.Article {
	return articleService.articleRepository.List(6)
}
