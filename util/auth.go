package util

import (
	"mini-pos/dto"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetAuthClaims(ctx echo.Context) *dto.UserClaims  {
	user := ctx.Get("user").(*jwt.Token)
	return user.Claims.(*dto.UserClaims)
}