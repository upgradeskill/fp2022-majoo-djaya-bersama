package usecase

import (
	"mini-pos/dto"
	"mini-pos/repository"
	"mini-pos/util"
	"time"

	"github.com/labstack/echo/v4"
)

type TransactionUseCase interface {
	Insert(ctx echo.Context) (dto.TransactionPayload, []dto.ValidationMessage, error)
	GetAll(ctx echo.Context) ([]dto.Transaction, error)
}

type transactionUseCase struct {
	transactionRepository repository.TransactionRepository
}

func InitTransactionUseCase(transactionRepository repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepository: transactionRepository,
	}
}

func (uc *transactionUseCase) Insert(ctx echo.Context) (data dto.TransactionPayload, invalidParameter []dto.ValidationMessage, err error) {

	if err = ctx.Bind(&data); err != nil {
		return
	}

	// validation
	// TODO: validate payload
	if data.OrderNumber == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "order_number", Message: "order_number is required"})
	}

	if len(invalidParameter) > 0 {
		return
	}

	// setup data
	data.OrderDate = time.Now()
	data.PaymentDate = time.Now()
	// TODO: get userID from login

	var transaction dto.Transaction
	if transaction, err = uc.transactionRepository.Insert(data.ToModel()); err != nil {
		return
	}

	for _, detail := range data.TransactionDetail {
		detail.TransactionID = transaction.Id
		if _, err = uc.transactionRepository.InsertDetail(detail.ToModel()); err != nil {
			return
		}
	}

	return

}

func (uc *transactionUseCase) GetAll(ctx echo.Context) (data []dto.Transaction, err error) {
	var filter dto.Transaction
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	session, _ := util.SessionStore.Get(ctx.Request(), util.SESSION_ID)
	filter.UserID, _ = session.Values["user_id"].(uint)

	pagination := dto.InitPagination()
	if err = ctx.Bind(&pagination); err != nil {
		return
	}

	if data, err = uc.transactionRepository.GetAll(filter, pagination); err != nil {
		return
	}

	return
}
