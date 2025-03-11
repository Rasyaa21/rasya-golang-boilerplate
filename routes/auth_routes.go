package routes

import (
	"rasya-golang-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	api := router.Group("/auth")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}
}
