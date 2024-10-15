package db

import (
	"log"

	"github.com/TalesPalma/PerdirComida/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDb() {

	config := &gorm.Config{}
	DB, err = gorm.Open(sqlite.Open("database.db"), config)

	if err != nil {
		log.Println("Erro with connection database:", err)
	}

	err = DB.AutoMigrate(&models.Pedidos{})

	if err != nil {
		log.Println("Erro with migrate database:", err)
	}

}
