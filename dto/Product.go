package dto

type Product struct {
	Model
	CategoryId  uint   `json:"category_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	IsActive    bool   `json:"is_active"`
}