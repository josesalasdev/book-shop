// Package controllers login.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/book-shop/src/models"
	"github.com/josesalasdev/book-shop/src/services"
)

type signInReponse struct {
	Token string `json:"token"`
}

type signInModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lm *signInModel) Validate() bool {
	var user models.User
	if err := models.DB.First(&user, "email = ?", lm.Email).Error; err != nil {
		return false
	}
	return user.ValidatePassword(lm.Password)
}

// SignIn Controller.
// CreateCategory godoc
// @Summary SignIn.
// @Description SignIn.
// @Accept  json
// @Produce  json
// @Param data body signInModel true "signIn Data"
// @Success 200 {object} signInReponse
// @Router /signin/ [post]
func SignIn(c *gin.Context) {
	serviceJWT := services.JWTAuthService()

	var input signInModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !input.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email or Password invalid."})
		return
	}
	token := serviceJWT.GenerateToken(input.Email, true)
	c.JSON(http.StatusOK, signInReponse{token})
}
