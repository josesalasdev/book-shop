package models

import "time"

// ModelBase of fields for models.
type ModelBase struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
	DeleteAt  time.Time `json:"deleteAt" gorm:"index"`
}
