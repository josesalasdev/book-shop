// controllers/book.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/golang_api_template/src/models"
)

// GET /books
// Get all books
// Ping godoc
// @Summary Find Books.
// @Description Get all books.
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Book
// @Router /books [get]
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
