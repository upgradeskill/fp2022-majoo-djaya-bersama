package api

import (
	"mini-pos/repository"
	"mini-pos/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OutletHandler struct {
	outletUseCase usecase.OutletUseCase
}

func NewOutletHandler() *OutletHandler {
	outletRepo := repository.InitOutletRepository()
	outletUseCase := usecase.InitOutletUseCase(outletRepo)
	return &OutletHandler{
		outletUseCase: outletUseCase,
	}
}

func OutletApi(e *echo.Group) {
	outletHandler := NewOutletHandler()
	e.GET("/outlet", outletHandler.OutletGetAll)
	e.GET("/outlet/:ID", outletHandler.OutletGetById)
	e.POST("/outlet", outletHandler.OutletInsert)
	e.PUT("/outlet", outletHandler.OutletUpdate)
	e.DELETE("/outlet/:ID", outletHandler.OutletDelete)

}

func (hand *OutletHandler) OutletGetAll(c echo.Context) error {
	data, err := hand.outletUseCase.GetAll(c)
	return SetupResponseGet(c, data, err)
}

func (hand *OutletHandler) OutletGetById(c echo.Context) error {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, err := hand.outletUseCase.GetByID(c, uint(id))

	return SetupResponseGet(c, data, err)
}

func (hand *OutletHandler) OutletInsert(c echo.Context) error {
	data, validate, err := hand.outletUseCase.Insert(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *OutletHandler) OutletUpdate(c echo.Context) error {
	data, validate, err := hand.outletUseCase.Update(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *OutletHandler) OutletDelete(c echo.Context) error {
	resp := make(map[string]interface{})
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, validate, err := hand.outletUseCase.Delete(uint(id))
	return SetupResponsePost(c, data, validate, err)
}
