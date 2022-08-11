package api

import (
	"mini-pos/repository"
	"mini-pos/security"
	"mini-pos/usecase"
	"mini-pos/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler() *ProductHandler {
	productRepo := repository.InitProductRepository()
	authorizeRepo := repository.InitAuthorizeRepository()
	jwtService := security.JWTAuthService()
	productUseCase := usecase.InitProductUseCase(productRepo, authorizeRepo, jwtService)
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func ProductApi(e *echo.Group) {
	productHandler := NewProductHandler()
	e.GET("/products", productHandler.ProductList)
	e.GET("/products/:ID", productHandler.ProductShow)
	e.POST("/products", productHandler.ProductInsert)
	e.PUT("/products/:ID", productHandler.ProductUpdate)
	e.DELETE("/products/:ID", productHandler.ProductDelete)
}

func (ph *ProductHandler) ProductList(ctx echo.Context) error  {
	resp := make(map[string]interface{})
	products, validate, err := ph.productUseCase.ProductList(ctx, util.GetAuthClaims(ctx))
	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return ctx.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return ctx.JSON(http.StatusNotFound, resp)
	}
	resp["message"] = "Success"
	resp["data"] = products

	return ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) ProductShow(ctx echo.Context) error  {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	product, validate, err := ph.productUseCase.ProductShow(uint(id), util.GetAuthClaims(ctx))
	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return ctx.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return ctx.JSON(http.StatusNotFound, resp)
	}
	resp["message"] = "Success"
	resp["data"] = product
	return ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) ProductInsert(ctx echo.Context) error  {
	resp := make(map[string]interface{})
	product, validate, err := ph.productUseCase.ProductInsert(ctx, util.GetAuthClaims(ctx))
	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return ctx.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = product
	return ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) ProductUpdate(ctx echo.Context) error  {
	resp := make(map[string]interface{})
	product, validate, err := ph.productUseCase.ProductUpdate(ctx, util.GetAuthClaims(ctx))
	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return ctx.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = product
	return ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) ProductDelete(ctx echo.Context) error  {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	product, validate, err := ph.productUseCase.ProductDelete(uint(id), util.GetAuthClaims(ctx))
	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return ctx.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = product
	return ctx.JSON(http.StatusOK, resp)
}