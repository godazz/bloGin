package repositories

import (
	articleModel "github.com/godazz/bloGin/internal/modules/article/models"
	"github.com/godazz/bloGin/pkg/db"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: db.Connecion(),
	}
}

func (ArticleRepository *ArticleRepository) List(limit int) []articleModel.Article {
	var articles []articleModel.Article

	ArticleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}
