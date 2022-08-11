package usecase

import (
	"errors"
	"log"
	"mini-pos/dto"
	"mini-pos/repository"
	"mini-pos/security"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	ProductList(ctx echo.Context, claims *dto.UserClaims) ([]dto.Product, []dto.ValidationMessage, error)
	ProductShow(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
	ProductInsert(ctx echo.Context, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
	ProductUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
	ProductDelete(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error)
}

type OutletProductUseCase interface {
	OutletProductList(ctx echo.Context, claims *dto.UserClaims) ([]dto.OutletProduct, []dto.ValidationMessage, error)
	OutletProductShow(OutletId uint, ProductId uint, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error)
	OutletProductInsert(ctx echo.Context, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error)
	OutletProductUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error)
	OutletProductDelete(OutletId uint, ProductId uint, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error)
}

type productUseCase struct {
	productRepository   repository.ProductRepository
	authorizeRepository repository.AuthorizeRepository
	jwtService          security.JWTService
}

type outletProductUseCase struct {
	outletProductRepository   repository.OutletProductRepository
	authorizeRepository repository.AuthorizeRepository
	jwtService          security.JWTService
}

func InitProductUseCase(productRepository repository.ProductRepository, authorizeRepository repository.AuthorizeRepository, jwtService security.JWTService) ProductUseCase {
	return &productUseCase{
		productRepository:   productRepository,
		authorizeRepository: authorizeRepository,
		jwtService:          jwtService,
	}
}

func InitOutletProductUseCase(outletProductRepository repository.OutletProductRepository, authorizeRepository repository.AuthorizeRepository, jwtService security.JWTService) OutletProductUseCase {
	return &outletProductUseCase{
		outletProductRepository: outletProductRepository,
		authorizeRepository    : authorizeRepository,
		jwtService             : jwtService,
	}
}

func (pc *productUseCase) ProductList(ctx echo.Context, claims *dto.UserClaims) (products []dto.Product, error_validation []dto.ValidationMessage, err error) {
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

	var filter dto.Product
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	pagination := dto.InitPagination()
	if err = ctx.Bind(&pagination); err != nil {
		return
	}

	products, err = pc.productRepository.List(filter, pagination)
	if err != nil {
		return []dto.Product{}, nil, err
	}
	return products, nil, nil
}

func (pc *productUseCase) ProductShow(id uint, claims *dto.UserClaims) (dto.Product, []dto.ValidationMessage, error) {
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

	payload.IsActive = dto.IsActive{IsActive: 1}
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

	id, err := strconv.Atoi(ctx.Param("ID"))
	old, err := pc.productRepository.Show(uint(id))

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to get data")
	}

	payload.Id = old.Id

	//validation
	if payload.CategoryId == 0 {
		payload.CategoryId = old.CategoryId
	}

	if payload.Name == "" {
		payload.Name = old.Name
	}

	if payload.Description == "" {
		payload.Description = old.Description
	}

	if payload.ImagePath == "" {
		payload.ImagePath = old.ImagePath
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

// ========================== Start Outlet Product ========================== 

func (opc *outletProductUseCase) OutletProductList(ctx echo.Context, claims *dto.UserClaims) (products []dto.OutletProduct, error_validation []dto.ValidationMessage, err error) {
	OutletId, err := strconv.Atoi(ctx.Param("OutletId"))
	if err != nil {
		return []dto.OutletProduct{}, nil, errors.New("outlet_id is required")
	}
	
	if claims.Role == 1 {
		err := opc.authorizeRepository.OwnerAuthorize(claims.Id, uint(OutletId))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return []dto.OutletProduct{}, nil, errors.New("you don't have right")
			}
			// return []dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := opc.authorizeRepository.StaffAuthorize(claims.Id, uint(OutletId))
		if err != nil {
			// return []dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	var filter dto.OutletProduct
	if err = ctx.Bind(&filter); err != nil {
		return
	}

	pagination := dto.InitPagination()
	if err = ctx.Bind(&pagination); err != nil {
		return
	}

	products, err = opc.outletProductRepository.List(filter, pagination)
	if err != nil {
		return []dto.OutletProduct{}, nil, err
	}
	return products, nil, nil
}

func (opc *outletProductUseCase) OutletProductShow(OutletId uint, ProductId uint, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error) {
	if claims.Role == 1 {
		err := opc.authorizeRepository.OwnerAuthorize(claims.Id, uint(1))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.OutletProduct{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := opc.authorizeRepository.StaffAuthorize(claims.Id, uint(1))
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}
	response, err := opc.outletProductRepository.Show(OutletId, ProductId)
	if err != nil {
		return dto.OutletProduct{}, nil, err
	}
	return response, nil, nil
}

func (opc *outletProductUseCase) OutletProductInsert(ctx echo.Context, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error) {

	payload := dto.OutletProduct{}

	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}

	//validation
	if payload.OutletID == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "outlet_id", Message: "outlet_id is required"})
	}

	if payload.ProductID == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "product_id", Message: "product_id is required"})
	}

	if decimal.Zero.Equal(payload.Price) {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "Price", Message: "Price is required"})
	}

	if payload.Stock == uint(0) {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "stock", Message: "stock is required"})
	}

	if claims.Role == 1 {
		err := opc.authorizeRepository.OwnerAuthorize(claims.Id, payload.OutletID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.OutletProduct{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := opc.authorizeRepository.StaffAuthorize(claims.Id, payload.OutletID)
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	if len(invalidParameter) > 0 {
		return dto.OutletProduct{}, invalidParameter, nil
	}

	payload.IsActive = dto.IsActive{IsActive: 1}
	payload.CreatedBy = claims.Id

	response, err := opc.outletProductRepository.Insert(payload)
	if err != nil {
		return dto.OutletProduct{}, nil, err
	}
	return response, nil, nil
}

func (opc *outletProductUseCase) OutletProductUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error) {
	payload := dto.OutletProduct{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}

	//validation
	if payload.OutletID <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "outlet_id", Message: "outlet_id is required"})
	}

	if payload.ProductID <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "product_id", Message: "product_id is required"})
	}

	if claims.Role == 1 {
		err := opc.authorizeRepository.OwnerAuthorize(claims.Id, payload.OutletID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.OutletProduct{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := opc.authorizeRepository.StaffAuthorize(claims.Id, payload.OutletID)
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	if len(invalidParameter) > 0 {
		return dto.OutletProduct{}, invalidParameter, nil
	}

	payload.IsActive = dto.IsActive{IsActive: 1}
	payload.UpdatedBy = claims.Id

	response, err := opc.outletProductRepository.Update(payload)
	if err != nil {
		return dto.OutletProduct{}, nil, err
	}
	return response, nil, nil
}

func (opc *outletProductUseCase) OutletProductDelete(OutletId uint, ProductId uint, claims *dto.UserClaims) (dto.OutletProduct, []dto.ValidationMessage, error) {
	var invalidParameter []dto.ValidationMessage
	product := dto.OutletProduct{}

	//validation
	if OutletId <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "outlet_id", Message: "outlet_id is required"})
	}

	if ProductId <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "product_id", Message: "product_id is required"})
	}

	if claims.Role == 1 {
		err := opc.authorizeRepository.OwnerAuthorize(claims.Id, OutletId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return dto.OutletProduct{}, nil, errors.New("you don't have right")
			}
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}

	} else {
		err := opc.authorizeRepository.StaffAuthorize(claims.Id, OutletId)
		if err != nil {
			// return dto.Product{}, nil, errors.New("failed to authorize")
		}
	}

	if len(invalidParameter) > 0 {
		return product, invalidParameter, nil
	}
	response, err := opc.outletProductRepository.DeleteByID(OutletId, ProductId)
	if err != nil {
		return product, nil, err
	}
	return response, nil, nil
}
