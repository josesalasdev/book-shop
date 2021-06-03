//models.category

package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CategoryInput struct {
	Name string `json:"name"`
}
