package responses

import (
	"fmt"

	articleModel "github.com/godazz/bloGin/internal/modules/article/models"
	"github.com/godazz/bloGin/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      responses.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     "/assets/img/demopic/10.jpg",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      responses.ToUser(article.User),
	}
}

func ToArticles(articles []articleModel.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
