package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	NameFood      string `json:"name" gorm:"not null"`
	Price         string `json:"price" gorm:"not null"`
	Calories      string `json:"calories"`
	Proteins      string `json:"proteins"`
	Fats          string `json:"fats"`
	Carbohydrates string `json:"carbohydrates"`
	Sodium        string `json:"sodium"`
	Author        Author `json:"author" gorm:"embedded"`
}
