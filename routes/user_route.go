package routes

import (
	"rasya-golang-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	api := router.Group("user")
	{
		api.GET("/", controllers.CurrentUser)
	}
}
