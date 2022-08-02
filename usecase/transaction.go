package usecase

import (
	"mini-pos/dto"
	"mini-pos/repository"
	"time"

	"github.com/labstack/echo/v4"
)

type TransactionUseCase interface {
	Insert(ctx echo.Context) (dto.TransactionPayload, []dto.ValidationMessage, error)
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
	if data.OrderNumber == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "order_number", Message: "order_number is required"})
	}

	if len(invalidParameter) > 0 {
		return
	}

	// setup data
	data.OrderDate = time.Now()
	data.PaymentDate = time.Now()

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
