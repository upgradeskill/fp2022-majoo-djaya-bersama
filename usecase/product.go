package usecase

import (
	"errors"
	"log"
	"mini-pos/dto"
	"mini-pos/repository"
	"mini-pos/security"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	ProductList(page int, pageSize int, claims *dto.UserClaims) ([]dto.Product, []dto.ValidationMessage, error)
	ProductShow(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
	ProductInsert(ctx echo.Context, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
	ProductUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
	ProductDelete(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
}

type productUseCase struct {
	productRepository      repository.ProductRepository
	authorizeRepository repository.AuthorizeRepository
	jwtService          security.JWTService
}

func InitProductUseCase(productRepository repository.ProductRepository, authorizeRepository repository.AuthorizeRepository, jwtService security.JWTService) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
		authorizeRepository: authorizeRepository,
		jwtService: jwtService,
	}
}

func (pc *productUseCase) ProductList(page int, pageSize int, claims *dto.UserClaims) ([]dto.Product, []dto.ValidationMessage, error)  {
	if claims.Role == 1 {
		err := pc.authorizeRepository.OwnerAuthorize(claims.Id, uint(1))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return []dto.Product{}, nil, errors.New("you don't have right")
			}
			// return []dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := pc.authorizeRepository.StaffAuthorize(claims.Id, uint(1))
		if err != nil {
			// return []dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	products, err := pc.productRepository.List(page, pageSize)
	if err != nil {
		return []dto.Product{}, nil, err
	}
	return products, nil, nil
}

func (pc *productUseCase) ProductShow(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)  {
	if claims.Role == 1 {
		err := pc.authorizeRepository.OwnerAuthorize(claims.Id, uint(1))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.Product{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := pc.authorizeRepository.StaffAuthorize(claims.Id, uint(1))
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}
	response, err := pc.productRepository.Show(id)
	if err != nil {
		return dto.Product{}, nil, err
	}
	return response, nil, nil
}

func (pc *productUseCase) ProductInsert(ctx echo.Context, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error) {

	payload := dto.Product{}

	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}

	//validation
	if payload.CategoryId == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "category_id", Message: "category_id is required"})
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.Description == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "description", Message: "description is required"})
	}

	if payload.ImagePath == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "image path", Message: "image path is required"})
	}

	if claims.Role == 1 {
		err := pc.authorizeRepository.OwnerAuthorize(claims.Id, uint(1))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.Product{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := pc.authorizeRepository.StaffAuthorize(claims.Id, uint(1))
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	if len(invalidParameter) > 0 {
		return dto.Product{}, invalidParameter, nil
	}

	payload.IsActive = true
	payload.CreatedBy = claims.Id

	response, err := pc.productRepository.Insert(payload)
	if err != nil {
		return dto.Product{}, nil, err
	}
	return response, nil, nil
}

func (pc *productUseCase) ProductUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error) {
	payload := dto.Product{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}
	
	//validation
	if payload.Id <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}
	
	if payload.CategoryId == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "category_id", Message: "category_id is required"})
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.Description == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "description", Message: "description is required"})
	}

	if payload.ImagePath == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "image path", Message: "image path is required"})
	}

	if claims.Role == 1 {
		err := pc.authorizeRepository.OwnerAuthorize(claims.Id, uint(1))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.Product{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := pc.authorizeRepository.StaffAuthorize(claims.Id, uint(1))
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	if len(invalidParameter) > 0 {
		return dto.Product{}, invalidParameter, nil
	}

	payload.IsActive = true
	payload.UpdatedBy = claims.Id

	response, err := pc.productRepository.Update(payload)
	if err != nil {
		return dto.Product{}, nil, err
	}
	return response, nil, nil
}

func (pc *productUseCase) ProductDelete(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error) {
	var invalidParameter []dto.ValidationMessage
	product := dto.Product{}
	
	//validation
	if id == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}

	if claims.Role == 1 {
		err := pc.authorizeRepository.OwnerAuthorize(claims.Id, uint(1))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.Product{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := pc.authorizeRepository.StaffAuthorize(claims.Id, uint(1))
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	if len(invalidParameter) > 0 {
		return product, invalidParameter, nil
	}
	response, err := pc.productRepository.DeleteByID(id)
	if err != nil {
		return product, nil, err
	}
	return response, nil, nil
}