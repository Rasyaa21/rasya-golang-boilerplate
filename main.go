package main

import (
	"fmt"
	"rasya-golang-boilerplate/config"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer config.DisconnectDB(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.Run(":8080")
}
