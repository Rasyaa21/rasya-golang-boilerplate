package middlewares

import (
	"net/http"

	"rasya-golang-boilerplate/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "unathorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
