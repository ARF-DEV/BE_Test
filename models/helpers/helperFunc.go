package helpers

import (
	"github.com/ARF-DEV/BE_Test/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {

	cat := []models.Category{
		{Name: "Sayuran", Status: true, Image: "images/popo.png"},
	}

	for _, item := range cat {
		db.Create(&item)
	}

	var sayuran models.Category

	db.First(&sayuran, "Name = ?", "Sayuran")

	products := []models.Product{
		{
			Name: "Wortel", Description: "Ini wortel", Stock: true, Price: 10000,
			Price_type: 500, ProductImages: []models.ProductImage{
				{ImagePath: "images/wortel1.png"},
				{ImagePath: "images/wortel2.png"}},
			Status: true, Category: sayuran,
		},
		{
			Name: "Kentang", Description: "Ini Kentang", Stock: true, Price: 10000,
			Price_type: 500, ProductImages: []models.ProductImage{
				{ImagePath: "images/Kentang1.png"},
				{ImagePath: "images/Kentang2.png"}},
			Status: true, Category: sayuran,
		},
	}
	for _, item := range products {
		db.Create(&item)
	}

}
