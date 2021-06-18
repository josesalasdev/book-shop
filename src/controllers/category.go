// controllers/category.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/golang_api_template/src/models"
)

// CreateCategory godoc
// @Summary Create a category.
// @Description Create a category.
// @Accept  json
// @Produce  json
// @Param data body models.CategoryInput true "Category Data"
// @Success 200 {object} models.Category
// @Router /category/ [post]
func CreateCategory(c *gin.Context) {
	// Validate input
	var input models.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create category
	category := models.Category{Name: input.Name}
	models.DB.Create(&category)

	c.JSON(http.StatusCreated, input)
}

// FindCategory godoc
// @Summary List categories.
// @Description Create a category.
// @Accept  json
// @Produce  json
// @Param data body models.Category true "Category Data"
// @Success 200 {object} []models.Category
// @Router /category/ [get]
func FindCategory(c *gin.Context) {
	var categories []models.Category
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}
