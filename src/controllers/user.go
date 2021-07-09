// controllers/book.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/golang_api_template/src/models"
)

// CreateUser controller.
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}
	user.EncryptPassword()
	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}
