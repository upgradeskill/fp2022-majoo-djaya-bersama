package api

import (
	"mini-pos/repository"
	"mini-pos/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionHandler() *TransactionHandler {
	transactionRepo := repository.InitTransactionRepository()
	transactionUseCase := usecase.InitTransactionUseCase(transactionRepo)
	return &TransactionHandler{
		transactionUseCase: transactionUseCase,
	}
}

func TransactionApi(e *echo.Group) {
	transactionHandler := NewTransactionHandler()
	e.GET("/transaction", transactionHandler.TransactionGet)
	e.POST("/transaction", transactionHandler.TransactionInsert)
}

func (hand *TransactionHandler) TransactionInsert(c echo.Context) error {
	resp := make(map[string]interface{})
	data, validate, err := hand.transactionUseCase.Insert(c)

	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}

func (hand *TransactionHandler) TransactionGet(c echo.Context) error {
	resp := make(map[string]interface{})
	data, err := hand.transactionUseCase.GetAll(c)

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}
