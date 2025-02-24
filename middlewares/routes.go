package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrintRoutes(router *gin.Engine) {
	fmt.Println("Registered Routes:")
	for _, route := range router.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
}
