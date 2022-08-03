package dto

type Category struct {
	Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	IsActive
}
