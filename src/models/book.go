//models.book

package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title      string  `json:"title"`
	Price      float32 `json:"price"`
	Author     string  `json:"author"`
	CategoryID uint    `json:"category"`
}

type CreateBookInput struct {
	Title      string  `json:"title" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	Author     string  `json:"author" binding:"required"`
	CategoryID uint    `json:"category"`
}
