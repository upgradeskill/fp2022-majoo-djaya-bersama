package usecase

import (
	"errors"
	"log"
	"mini-pos/dto"
	"mini-pos/repository"
	"mini-pos/security"

	"github.com/labstack/echo/v4"
)

type OutletUseCase interface {
	GetAll(echo.Context) ([]dto.OutletResponse, error)
	GetByID(echo.Context, uint) (dto.OutletResponse, error)
	Insert(ctx echo.Context, claims *dto.UserClaims) (dto.Outlet, []dto.ValidationMessage, error)
	Update(ctx echo.Context, claims *dto.UserClaims) (dto.Outlet, []dto.ValidationMessage, error)
	Delete(id uint) (dto.Outlet, []dto.ValidationMessage, error)
}

type outletUseCase struct {
	outletRepository    repository.OutletRepository
	authorizeRepository repository.AuthorizeRepository
	jwtService          security.JWTService
}

func InitOutletUseCase(outletRepository repository.OutletRepository, authorizeRepository repository.AuthorizeRepository, jwtService security.JWTService) OutletUseCase {
	return &outletUseCase{
		outletRepository:    outletRepository,
		authorizeRepository: authorizeRepository,
		jwtService:          jwtService,
	}
}

func (uc *outletUseCase) GetAll(ctx echo.Context) (data []dto.OutletResponse, err error) {
	var filter dto.Outlet
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	pagination := dto.InitPagination()
	if err = ctx.Bind(&pagination); err != nil {
		return
	}

	var outlets []dto.Outlet
	if outlets, err = uc.outletRepository.GetAll(filter, pagination); err != nil {
		return
	}

	for _, outlet := range outlets {
		data = append(data, dto.OutletResponse{
			Id:       outlet.Id,
			Name:     outlet.Name,
			Address:  outlet.Address,
			IsActive: outlet.IsActive.IsActive,
		})
	}
	return
}

func (uc *outletUseCase) GetByID(ctx echo.Context, id uint) (data dto.OutletResponse, err error) {
	var filter dto.Outlet
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	filter.Id = id

	var outlet dto.Outlet
	if outlet, err = uc.outletRepository.GetByID(id); err != nil {
		return
	}
	data = dto.OutletResponse{
		Id:       outlet.Id,
		Name:     outlet.Name,
		Address:  outlet.Address,
		IsActive: outlet.IsActive.IsActive,
	}
	return
}

func (uc *outletUseCase) Insert(ctx echo.Context, claims *dto.UserClaims) (dto.Outlet, []dto.ValidationMessage, error) {
	payload := dto.Outlet{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.Address == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "address", Message: "address is required"})
	}

	if len(invalidParameter) > 0 {
		return dto.Outlet{}, invalidParameter, nil
	}

	payload.CreatedBy = claims.Id

	response, err := uc.outletRepository.Insert(payload)
	if err != nil {
		return dto.Outlet{}, nil, err
	}
	return response, nil, nil
}

func (uc *outletUseCase) Update(ctx echo.Context, claims *dto.UserClaims) (dto.Outlet, []dto.ValidationMessage, error) {
	payload := dto.Outlet{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}

	if payload.Id <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.Address == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "address", Message: "address is required"})
	}

	if len(invalidParameter) > 0 {
		return dto.Outlet{}, invalidParameter, nil
	}

	payload.UpdatedBy = claims.Id

	response, err := uc.outletRepository.Update(payload)
	if err != nil {
		return dto.Outlet{}, nil, err
	}
	return response, nil, nil
}

func (uc *outletUseCase) Delete(id uint) (dto.Outlet, []dto.ValidationMessage, error) {
	var invalidParameter []dto.ValidationMessage
	user := dto.Outlet{}

	if id == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}

	if len(invalidParameter) > 0 {
		return user, invalidParameter, nil
	}
	response, err := uc.outletRepository.DeleteByID(id)
	if err != nil {
		return user, nil, err
	}
	return response, nil, nil
}
