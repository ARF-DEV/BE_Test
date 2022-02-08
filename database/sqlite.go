package database

import (
	"fmt"

	"github.com/ARF-DEV/BE_Test/models"
	"github.com/ARF-DEV/BE_Test/models/helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup(db *gorm.DB) {
	db.AutoMigrate(&models.Product{}, &models.Category{}, &models.ProductImage{})
	helpers.Seed(db)
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("entity.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected to DB")

	Setup(DB)
}
