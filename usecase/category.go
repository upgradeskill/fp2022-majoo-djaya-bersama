package usecase

import (
	"errors"
	"log"
	"mini-pos/dto"
	"mini-pos/repository"

	"github.com/labstack/echo/v4"
)

type CategoryUseCase interface {
	GetAll(echo.Context) ([]dto.CategoryResponse, error)
	GetByID(echo.Context, uint) (dto.CategoryResponse, error)
	Insert(ctx echo.Context) (dto.Category, []dto.ValidationMessage, error)
	Update(ctx echo.Context) (dto.Category, []dto.ValidationMessage, error)
	Delete(id uint) (dto.Category, []dto.ValidationMessage, error)
}

type categoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

func InitCategoryUseCase(categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *categoryUseCase) GetAll(ctx echo.Context) (data []dto.CategoryResponse, err error) {
	var filter dto.Category
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	pagination := dto.InitPagination()
	if err = ctx.Bind(&pagination); err != nil {
		return
	}

	var categories []dto.Category
	if categories, err = uc.categoryRepository.GetAll(filter, pagination); err != nil {
		return
	}

	for _, category := range categories {
		data = append(data, dto.CategoryResponse{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive.IsActive,
		})
	}
	return
}

func (uc *categoryUseCase) GetByID(ctx echo.Context, id uint) (data dto.CategoryResponse, err error) {
	var filter dto.Category
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	filter.Id = id

	var category dto.Category
	if category, err = uc.categoryRepository.GetByID(id); err != nil {
		return
	}
	data = dto.CategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive.IsActive,
	}
	return
}

func (uc *categoryUseCase) Insert(ctx echo.Context) (dto.Category, []dto.ValidationMessage, error) {
	payload := dto.Category{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.Description == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "address", Message: "address is required"})
	}

	if len(invalidParameter) > 0 {
		return dto.Category{}, invalidParameter, nil
	}

	payload.IsActive = dto.IsActive{IsActive: 1}

	response, err := uc.categoryRepository.Insert(payload)
	if err != nil {
		return dto.Category{}, nil, err
	}
	return response, nil, nil
}

func (uc *categoryUseCase) Update(ctx echo.Context) (dto.Category, []dto.ValidationMessage, error) {
	payload := dto.Category{}
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

	if payload.Description == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "address", Message: "address is required"})
	}

	if len(invalidParameter) > 0 {
		return dto.Category{}, invalidParameter, nil
	}

	payload.IsActive = dto.IsActive{IsActive: 1}

	response, err := uc.categoryRepository.Update(payload)
	if err != nil {
		return dto.Category{}, nil, err
	}
	return response, nil, nil
}

func (uc *categoryUseCase) Delete(id uint) (dto.Category, []dto.ValidationMessage, error) {
	var invalidParameter []dto.ValidationMessage
	user := dto.Category{}

	if id == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}

	if len(invalidParameter) > 0 {
		return user, invalidParameter, nil
	}
	response, err := uc.categoryRepository.DeleteByID(id)
	if err != nil {
		return user, nil, err
	}
	return response, nil, nil
}
