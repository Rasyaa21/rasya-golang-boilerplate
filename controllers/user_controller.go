package controllers

import (
	"fmt"
	"net/http"
	"rasya-golang-boilerplate/models"
	"rasya-golang-boilerplate/utils/token"
	"rasya-golang-boilerplate/validation"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input validation.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.StoreUser(input.Name, input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	token, _ := token.GenerateToken(user.ID)
	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{"message": token})
}

func Login(c *gin.Context) {
	var input validation.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	isValid, err := user.CheckPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking password"})
		return
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := token.GenerateToken(user.ID)
	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{"message": token})
}

func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
