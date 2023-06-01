package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ArticleService "github.com/godazz/bloGin/internal/modules/article/services"
	"github.com/godazz/bloGin/pkg/html"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}
