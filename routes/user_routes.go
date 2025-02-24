package routes

import (
	"rasya-golang-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	api := router.Group("/users")
	{
		api.GET("/", controllers.GetAllUsers)
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}
}
