package repositories

import articleModel "github.com/godazz/bloGin/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []articleModel.Article
}
