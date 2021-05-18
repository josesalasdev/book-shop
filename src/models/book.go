//models.book

package models

type Book struct {
	ID    uint    `json: "id" gorm: "primary_key"`
	Title string  `json: "title"`
	Price float32 `json: "price"`
}
