package dto

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	Model
	OutletId    uint   `json:"outlet_id,omitempty"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	IsRole      int    `json:"is_role"`
	IsActive    int    `json:"is_active"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	IsRole   int    `json:"is_role"`
	IsActive int    `json:"is_active"`
	Token    string `json:"token"`
	// TokenRefresh string `json:"token_refresh"`
}

type UserClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
	jwt.StandardClaims
}
