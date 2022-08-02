package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Model
	OutletID       uint            `json:"outlet_id" form:"outlet_id"`
	Outlet         Outlet          `json:"outlet" form:"outlet" gorm:"foreignKey:OutletID"`
	UserID         uint            `json:"user_id" form:"user_id"`
	User           User            `json:"user" form:"user" gorm:"foreignKey:UserID"`
	OrderNumber    string          `json:"order_number" form:"order_number" binding:"required"`
	OrderDate      time.Time       `json:"order_date" form:"order_date" binding:"required"`
	OrderNominal   decimal.Decimal `json:"order_nominal" form:"order_nominal" gorm:"type:decimal(12,2)" binding:"required"`
	PaymentNumber  string          `json:"payment_number" form:"payment_number"`
	PaymentDate    time.Time       `json:"payment_date" form:"payment_date"`
	PaymentNominal decimal.Decimal `json:"payment_nominal" form:"payment_nominal" gorm:"type:decimal(12,2)"`
	PaymentMethod  uint            `json:"payment_method" form:"payment_method"`
	PaymentNote    string          `json:"payment_note" form:"payment_note"`
	IsStatus       uint            `json:"is_status" form:"is_status"`
}

type TransactionDetail struct {
	Model
	TransactionID uint            `json:"transaction_id" form:"transaction_id"`
	Transaction   Transaction     `json:"transaction" form:"transaction" gorm:"foreignKey:TransactionID"`
	ProductID     uint            `json:"product_id" form:"product_id"`
	Product       Product         `json:"product" form:"product" gorm:"foreignKey:ProductID"`
	ProductName   string          `json:"product_name" form:"product_name"`
	Quantity      uint            `json:"quantity" form:"quantity"`
	Price         decimal.Decimal `json:"price" form:"price" gorm:"type:decimal(12,2)"`
}

type TransactionDetailPayload struct {
	TransactionID uint            `json:"transaction_id" form:"transaction_id"`
	ProductID     uint            `json:"product_id" form:"product_id"`
	Product       Product         `json:"product" form:"product" gorm:"foreignKey:ProductID"`
	ProductName   string          `json:"product_name" form:"product_name"`
	Quantity      uint            `json:"quantity" form:"quantity"`
	Price         decimal.Decimal `json:"price" form:"price" gorm:"type:decimal(12,2)"`
}

type TransactionPayload struct {
	TransactionID     uint                       `json:"transaction_id" form:"transaction_id"`
	OutletID          uint                       `json:"outlet_id" form:"outlet_id"`
	Outlet            Outlet                     `json:"outlet" form:"outlet" gorm:"foreignKey:OutletID"`
	UserID            uint                       `json:"user_id" form:"user_id"`
	User              User                       `json:"user" form:"user" gorm:"foreignKey:UserID"`
	OrderNumber       string                     `json:"order_number" form:"order_number" binding:"required"`
	OrderDate         time.Time                  `json:"order_date" form:"order_date" binding:"required"`
	OrderNominal      decimal.Decimal            `json:"order_nominal" form:"order_nominal" gorm:"type:decimal(12,2)" binding:"required"`
	PaymentNumber     string                     `json:"payment_number" form:"payment_number"`
	PaymentDate       time.Time                  `json:"payment_date" form:"payment_date"`
	PaymentNominal    decimal.Decimal            `json:"payment_nominal" form:"payment_nominal" gorm:"type:decimal(12,2)"`
	PaymentMethod     uint                       `json:"payment_method" form:"payment_method"`
	PaymentNote       string                     `json:"payment_note" form:"payment_note"`
	IsStatus          uint                       `json:"is_status" form:"is_status"`
	TransactionDetail []TransactionDetailPayload `json:"transaction_detail"`
}

func (payload *TransactionPayload) ToModel() Transaction {
	return Transaction{
		OutletID:       payload.OutletID,
		UserID:         payload.UserID,
		OrderNumber:    payload.OrderNumber,
		OrderDate:      payload.OrderDate,
		OrderNominal:   payload.OrderNominal,
		PaymentNumber:  payload.PaymentNumber,
		PaymentDate:    payload.PaymentDate,
		PaymentNominal: payload.PaymentNominal,
		PaymentMethod:  payload.PaymentMethod,
		PaymentNote:    payload.PaymentNote,
		IsStatus:       payload.IsStatus,
	}
}

func (payload *TransactionDetailPayload) ToModel() TransactionDetail {
	return TransactionDetail{
		TransactionID: payload.TransactionID,
		ProductID:     payload.ProductID,
		ProductName:   payload.ProductName,
		Quantity:      payload.Quantity,
		Price:         payload.Price,
	}
}
