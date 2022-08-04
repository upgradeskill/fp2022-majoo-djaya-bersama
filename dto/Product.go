package dto

import "github.com/shopspring/decimal"

type Product struct {
	Model
	CategoryId  uint     `json:"category_id" form:"category_id"`
	Category    Category `json:"category" form:"category" gorm:"foreignKey:CategoryId"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	ImagePath   string   `json:"image_path" form:"image_path"`
	IsActive
}

type OutletProduct struct {
	Model
	OutletID  uint            `json:"outlet_id" form:"outlet_id" gorm:"primaryKey;autoIncrement:false"`
	Outlet    Outlet          `json:"outlet" form:"outlet" gorm:"foreignKey:OutletID"`
	ProductID uint            `json:"product_id" form:"product_id" gorm:"primaryKey;autoIncrement:false"`
	Product   Product         `json:"product" form:"product" gorm:"foreignKey:ProductID"`
	Stock     uint            `json:"stock" form:"stock"`
	Price     decimal.Decimal `json:"price" form:"price" gorm:"type:decimal(10,2)"`
	IsActive
}
