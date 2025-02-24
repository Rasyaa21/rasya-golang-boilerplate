package controllers

import (
	"net/http"
	"rasya-golang-boilerplate/models"
	"rasya-golang-boilerplate/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "hello",
	})
}

func Register(c *gin.Context) {
	var input validation.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.StoreUser(input.Name, input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "validated"})
}

func Login(c *gin.Context) {
	var input validation.LoginInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(input.Email)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	isValid, err := user.CheckPassword(input.Password)
	if err != nil || !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	//implement the jwt
	c.JSON(http.StatusOK, gin.H{"message": "login successfull"})
}
