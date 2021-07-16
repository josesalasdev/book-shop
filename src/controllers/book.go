// controllers/book.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/book-shop/src/models"
)

// GET /books
// Get all books
// Ping godoc
// @Summary Find books.
// @Description Get all books.
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Book
// @Router /book/ [get]
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// DetailBooks godoc
// @Summary Detail of a book.
// @Description Detail of a book by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /book/{id}/ [get]
func DetailBooks(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary Create a book.
// @Description Create a book.
// @Accept  json
// @Produce  json
// @Param data body models.CreateBookInput true "Book Data"
// @Success 200 {object} models.CreateBookInput
// @Router /book/ [post]
func CreateBook(c *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{
		Title:      input.Title,
		Price:      input.Price,
		Author:     input.Author,
		CategoryID: input.CategoryID,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, input)
}
