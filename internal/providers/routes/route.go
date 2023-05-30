package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "github.com/godazz/bloGin/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}
