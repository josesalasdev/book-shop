// controllers/book.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/book-shop/src/models"
)

// userResponse is a response.
type userResponse struct {
	Message string `json:"message"`
}

// CreateUser controller.
// @Summary Create an user.
// @Description Create an user.
// @Accept  json
// @Produce  json
// @Param data body models.UserInput true "user Data"
// @Success 200 {object} userResponse
// @Router /user/ [post]
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

	c.JSON(http.StatusCreated, userResponse{"OK"})
}
