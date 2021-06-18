//models.category

package models

// Category model.
type Category struct {
	ModelBase
	Name  string `json:"name"`
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// CategoryInput for model Category create.
type CategoryInput struct {
	Name string `json:"name"`
}
