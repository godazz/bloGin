package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/godazz/bloGin/internal/modules/home/controllers"
)

func Routes(router *gin.Engine) {

	homeController := controllers.New()
	router.GET("/", homeController.Index)
}
