package models

import (
	"time"
)

type Product struct {
	ID            int `gorm:"primarykey"`
	Name          string
	Description   string
	Stock         bool
	Price         int
	Price_type    int
	ProductImages []ProductImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status        bool
	CategoryID    *int
	Category      Category `gorm:"constraint:OnUpdate:CASCADE,ONDELETE:SET NULL;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ProductImage struct {
	ID        int
	ImagePath string
	ProductID int
}

type Category struct {
	ID        int
	Name      string
	Status    bool
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
