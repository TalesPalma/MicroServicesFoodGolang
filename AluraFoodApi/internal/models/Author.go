package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	NameAuthor string `json:"name" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
}
