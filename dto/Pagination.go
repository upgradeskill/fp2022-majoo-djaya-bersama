package dto

import (
	"mini-pos/helper"

	"gorm.io/gorm"
)

type Pagination struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func InitPagination() Pagination {
	return Pagination{Page: helper.PAGE, Limit: helper.LIMIT}
}

func (pagination *Pagination) Apply(db *gorm.DB) *gorm.DB {
	return db.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit)
}
