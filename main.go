package main

import (
	"fmt"
	"rasya-golang-boilerplate/config"
	"rasya-golang-boilerplate/middlewares"
	"rasya-golang-boilerplate/routes"

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
	api := r.Group("/api")

	routes.UserRoutes(api)
	middlewares.PrintRoutes(r)
	r.Run(":8080")
}
