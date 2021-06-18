//models.book

package models

// Book model.
type Book struct {
	ModelBase
	Title      string  `json:"title"`
	Price      float32 `json:"price"`
	Author     string  `json:"author"`
	CategoryID uint    `json:"category"`
}

// CreateBookInput for model book create.
type CreateBookInput struct {
	Title      string  `json:"title" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	Author     string  `json:"author" binding:"required"`
	CategoryID uint    `json:"category"`
}
