package api

import (
	"mini-pos/repository"
	"mini-pos/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionHandler() *TransactionHandler {
	transactionRepo := repository.InitTransactionRepository()
	productRepo := repository.InitProductRepository()
	transactionUseCase := usecase.InitTransactionUseCase(transactionRepo, productRepo)
	return &TransactionHandler{
		transactionUseCase: transactionUseCase,
	}
}

func TransactionApi(e *echo.Group) {
	transactionHandler := NewTransactionHandler()
	e.GET("/transaction", transactionHandler.TransactionGet)
	e.GET("/transaction/:ID", transactionHandler.TransactionGetDetail)
	e.POST("/transaction", transactionHandler.TransactionInsert)
	e.PUT("/transaction", transactionHandler.TransactionUpdate)
	e.POST("/transaction/payment", transactionHandler.TransactionSavePayment)
	e.PUT("/transaction/:ID", transactionHandler.TransactionCancel)
	e.DELETE("/transaction/:ID", transactionHandler.TransactionDelete)

}

func (hand *TransactionHandler) TransactionInsert(c echo.Context) error {
	data, validate, err := hand.transactionUseCase.Insert(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *TransactionHandler) TransactionGet(c echo.Context) error {
	data, err := hand.transactionUseCase.GetAll(c)
	return SetupResponseGet(c, data, err)
}

func (hand *TransactionHandler) TransactionGetDetail(c echo.Context) error {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, err := hand.transactionUseCase.GetDetailByID(c, uint(id))

	return SetupResponseGet(c, data, err)
}

func (hand *TransactionHandler) TransactionUpdate(c echo.Context) error {
	data, validate, err := hand.transactionUseCase.Update(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *TransactionHandler) TransactionSavePayment(c echo.Context) error {
	data, validate, err := hand.transactionUseCase.SavePayment(c)
	return SetupResponsePost(c, data, validate, err)
}

func (hand *TransactionHandler) TransactionCancel(c echo.Context) error {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = hand.transactionUseCase.CancelTransaction(c, uint(id))

	return SetupResponseGet(c, nil, err)
}

func (hand *TransactionHandler) TransactionDelete(c echo.Context) error {
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = hand.transactionUseCase.Delete(c, uint(id))

	return SetupResponseGet(c, nil, err)
}
