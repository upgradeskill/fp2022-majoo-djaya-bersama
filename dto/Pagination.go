package dto

import (
	"mini-pos/util"

	"gorm.io/gorm"
)

type Pagination struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func InitPagination() Pagination {
	return Pagination{Page: util.PAGE, Limit: util.LIMIT}
}

func (pagination *Pagination) Apply(db *gorm.DB) *gorm.DB {
	return db.Offset(pagination.Page * pagination.Limit).Limit(pagination.Limit)
}
