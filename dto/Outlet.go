package dto

type Outlet struct {
	Model
	Name    string `json:"name"`
	Address string `json:"address"`
	IsActive
}
