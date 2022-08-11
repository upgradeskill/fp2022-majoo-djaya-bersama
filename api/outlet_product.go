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

type OutletProductHandler struct {
	outletProductUseCase usecase.OutletProductUseCase
}

func NewOutletProductHandler() *OutletProductHandler {
	outletProductRepo := repository.InitOutletProductRepository()
	authorizeRepo := repository.InitAuthorizeRepository()
	jwtService := security.JWTAuthService()
	outletProductUseCase := usecase.InitOutletProductUseCase(outletProductRepo, authorizeRepo, jwtService)
	return &OutletProductHandler{
		outletProductUseCase: outletProductUseCase,
	}
}

func OutletProductApi(e *echo.Group) {
	outletProductHandler := NewOutletProductHandler()
	e.GET("/outlets/:OutletId/products", outletProductHandler.OutletProductGetAll)
	e.GET("/outlets/:OutletId/products/:ProductId", outletProductHandler.OutletProductGet)
	e.POST("/outlets/:OutletId/products", outletProductHandler.OutletProductInsert)
	e.PUT("/outlets/:OutletId/products/:ProductId", outletProductHandler.OutletProductUpdate)
	e.DELETE("/outlets/:OutletId/products/:ProductId", outletProductHandler.OutletProductDelete)
}

func (oph *OutletProductHandler) OutletProductGetAll(ctx echo.Context) error  {
	resp := make(map[string]interface{})
	products, validate, err := oph.outletProductUseCase.OutletProductList(ctx, util.GetAuthClaims(ctx))
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

func (oph *OutletProductHandler) OutletProductGet(ctx echo.Context) error  {
	resp := make(map[string]interface{})

	OutletId, err := strconv.Atoi(ctx.Param("OutletId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	ProductId, err := strconv.Atoi(ctx.Param("ProductId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	product, validate, err := oph.outletProductUseCase.OutletProductShow(uint(OutletId), uint(ProductId), util.GetAuthClaims(ctx))
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

func (oph *OutletProductHandler) OutletProductInsert(ctx echo.Context) error  {
	resp := make(map[string]interface{})
	product, validate, err := oph.outletProductUseCase.OutletProductInsert(ctx, util.GetAuthClaims(ctx))
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

func (oph *OutletProductHandler) OutletProductUpdate(ctx echo.Context) error  {
	resp := make(map[string]interface{})
	product, validate, err := oph.outletProductUseCase.OutletProductUpdate(ctx, util.GetAuthClaims(ctx))
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

func (oph *OutletProductHandler) OutletProductDelete(ctx echo.Context) error  {
	resp := make(map[string]interface{})

	OutletId, err := strconv.Atoi(ctx.Param("OutletId"))
	if err != nil {
		resp["message"] = "invalid id"
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	ProductId, err := strconv.Atoi(ctx.Param("ProductId"))
	if err != nil {
		resp["message"] = "invalid id"
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	product, validate, err := oph.outletProductUseCase.OutletProductDelete(uint(OutletId), uint(ProductId), util.GetAuthClaims(ctx))
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