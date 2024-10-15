package db

import (
	"log"

	"github.com/TalesPalma/AluraFood/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabse() {
	configs := &gorm.Config{}
	DB, err = gorm.Open(sqlite.Open("Databse.db"), configs)
	if err != nil {
		log.Fatal("Connection with database failed:", err)
	}

	err = DB.AutoMigrate(&models.Food{})

	if err != nil {
		log.Fatal("Error with automigrate database : ", err)
	}

}
