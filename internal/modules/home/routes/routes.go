package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godazz/bloGin/pkg/html"
)

func Routes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home Page",
		})
	})
}
