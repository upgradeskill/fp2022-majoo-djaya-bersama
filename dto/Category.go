package dto

type Category struct {
	Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	IsActive
}

type CategoryResponse struct {
	Id          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	IsActive    int    `json:"is_active" form:"is_active"`
}
