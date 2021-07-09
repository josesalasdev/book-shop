// Package controllers login.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/golang_api_template/src/models"
	"github.com/josesalasdev/golang_api_template/src/services"
)

type loginReponse struct {
	Token string `json:"token"`
}

type loginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lm *loginModel) Validate() bool {
	var user models.User
	if err := models.DB.First(&user, "email = ?", lm.Email).Error; err != nil {
		return false
	}
	return user.ValidatePassword(lm.Password)
}

// SignIn Controller.
func SignIn(c *gin.Context) {
	serviceJWT := services.JWTAuthService()

	var input loginModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !input.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email or Password invalid."})
		return
	}
	token := serviceJWT.GenerateToken(input.Email, true)
	c.JSON(http.StatusOK, loginReponse{token})
}
