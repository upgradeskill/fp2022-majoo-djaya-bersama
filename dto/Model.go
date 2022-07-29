package dto

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint           `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at,omitempty" form:"created_at"`
	CreatedBy uint           `json:"created_by,omitempty" form:"created_by"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy uint           `json:"updated_by,omitempty" form:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" form:"deleted_at" gorm:"index,"`
	DeletedBy uint           `json:"deleted_by,omitempty" form:"deleted_by"`
}

type ValidationMessage struct {
	Parameter string `json:"parameter"`
	Message   string `json:"message"`
}
