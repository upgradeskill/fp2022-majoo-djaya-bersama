package dto

type Outlet struct {
	Model
	Name    string `json:"name"`
	Address string `json:"address"`
	IsActive
}

type OutletResponse struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	IsActive int    `json:"is_active" form:"is_active"`
}
