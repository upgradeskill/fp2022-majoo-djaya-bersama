package server

import (
	"log"
	"mini-pos/api"
	"mini-pos/dto"
	"mini-pos/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes() {
	e := echo.New()

	v1 := e.Group("/v1")

	// for public (unauthorized user)
	{
		api.Authorization(v1)
	}

	v1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &dto.UserClaims{},
		SigningKey:  []byte(util.GlobalConfig.JWT_SECRET),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}))

	// for authorized user
	{
		api.UserApi(v1)
		api.ProductApi(v1)
		api.OutletProductApi(v1)
		api.TransactionApi(v1)
		api.OutletApi(v1)
		api.CategoryApi(v1)
	}

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
