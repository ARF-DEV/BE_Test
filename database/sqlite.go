package database

import (
	"fmt"

	"github.com/ARF-DEV/BE_Test/models"
	"github.com/ARF-DEV/BE_Test/models/helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	DB.AutoMigrate(&models.Product{}, &models.Category{}, &models.ProductImage{})
	helpers.Seed(DB)
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("entity.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected to DB")

	Setup()
}
