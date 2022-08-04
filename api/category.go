package api

import (
	"mini-pos/repository"
	"mini-pos/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler() *CategoryHandler {
	categoryRepo := repository.InitCategoryRepository()
	categoryUseCase := usecase.InitCategoryUseCase(categoryRepo)
	return &CategoryHandler{
		categoryUseCase: categoryUseCase,
	}
}

func CategoryApi(e *echo.Group) {
	categoryHandler := NewCategoryHandler()
	e.GET("/category", categoryHandler.CategoryGetAll)
	e.GET("/category/:ID", categoryHandler.CategoryGetById)
	e.POST("/category", categoryHandler.CategoryInsert)
	e.PUT("/category", categoryHandler.CategoryUpdate)
	e.DELETE("/category/:ID", categoryHandler.CategoryDelete)

}

func (hand *CategoryHandler) CategoryGetAll(c echo.Context) error {
	data, err := hand.categoryUseCase.GetAll(c)
	return SetupResponseGet(c, data, err)
}

func (hand *CategoryHandler) CategoryGetById(c echo.Context) error {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, err := hand.categoryUseCase.GetByID(c, uint(id))

	return SetupResponseGet(c, data, err)
}

func (hand *CategoryHandler) CategoryInsert(c echo.Context) error {
	data, validate, err := hand.categoryUseCase.Insert(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *CategoryHandler) CategoryUpdate(c echo.Context) error {
	data, validate, err := hand.categoryUseCase.Update(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *CategoryHandler) CategoryDelete(c echo.Context) error {
	resp := make(map[string]interface{})
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, validate, err := hand.categoryUseCase.Delete(uint(id))
	return SetupResponsePost(c, data, validate, err)
}
