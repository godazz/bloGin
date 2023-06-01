package services

import (
	articleRepository "github.com/godazz/bloGin/internal/modules/article/repositories"
	articleResponse "github.com/godazz/bloGin/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: articleRepository.New(),
	}
}

func (articleService ArticleService) GetFeaturedArticles() articleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return articleResponse.ToArticles(articles)
}

func (articleService ArticleService) GetStoriesArticles() articleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return articleResponse.ToArticles(articles)
}
