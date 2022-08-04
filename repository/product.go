package repository

import (
	"mini-pos/database"
	"mini-pos/dto"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	List(filter dto.Product, pagination dto.Pagination) ([]dto.Product, error)
	Show(id uint) (dto.Product, error)
	Insert(dto.Product) (dto.Product, error)
	Update(dto.Product) (dto.Product, error)
	DeleteByID(ID uint) (data dto.Product, err error)
	GetOutletProductByID(ID int) (dto.OutletProduct, error)
}

type OutletProductRepository interface {
	List(filter dto.OutletProduct, pagination dto.Pagination) ([]dto.OutletProduct, error)
	Show(OutletId uint, ProductId uint) (dto.OutletProduct, error)
	Insert(dto.OutletProduct) (dto.OutletProduct, error)
	Update(dto.OutletProduct) (dto.OutletProduct, error)
	DeleteByID(OutletId uint, ProductId uint) (data dto.OutletProduct, err error)
}

type productRepo struct {
	DB *gorm.DB
}

type outletProductRepo struct {
	DB *gorm.DB
}

func InitProductRepository() ProductRepository {
	return &productRepo{
		DB: database.DB,
	}
}

func InitOutletProductRepository() OutletProductRepository {
	return &outletProductRepo{
		DB: database.DB,
	}
}

func (repo *productRepo) List(filter dto.Product, pagination dto.Pagination) (data []dto.Product, err error) {
	err = pagination.Apply(repo.DB).Find(&data, filter).Error
	return
}

func (repo *productRepo) Show(id uint) (data dto.Product, err error) {
	err = repo.DB.First(&data, id).Error
	return data, err
}

func (repo *productRepo) Insert(payload dto.Product) (data dto.Product, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *productRepo) Update(payload dto.Product) (data dto.Product, err error) {

	// get book by id
	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
		return
	}

	// update value
	data.CategoryId = payload.CategoryId
	data.Name = payload.Name
	data.Description = payload.Description
	data.ImagePath = payload.ImagePath
	data.IsActive = payload.IsActive
	data.UpdatedAt = time.Now()

	// update book data
	err = repo.DB.Save(&data).Error
	return
}

func (repo *productRepo) DeleteByID(ID uint) (data dto.Product, err error) {
	// get book by id
	if err = repo.DB.First(&data, ID).Error; err != nil {
		return
	}

	// delete book
	err = repo.DB.Delete(&data, ID).Error
	return
}

func (repo *productRepo) GetOutletProductByID(ID int) (data dto.OutletProduct, err error) {
	err = repo.DB.Preload("Product").Find(&data, ID).Error
	return
}

// ========================== Start Outlet Product ==========================

func (repo *outletProductRepo) List(filter dto.OutletProduct, pagination dto.Pagination) (data []dto.OutletProduct, err error) {
	err = pagination.Apply(repo.DB).Find(&data, filter).Error
	return
}

func (repo *outletProductRepo) Show(OutletId uint, ProductId uint) (data dto.OutletProduct, err error) {
	data = dto.OutletProduct{OutletID: OutletId, ProductID: ProductId}
	err = repo.DB.Preload("Product").First(&data).Where("outlet_id = ? and product_id = ?", OutletId, ProductId).Error
	return data, err
}

func (repo *outletProductRepo) Insert(payload dto.OutletProduct) (data dto.OutletProduct, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *outletProductRepo) Update(payload dto.OutletProduct) (data dto.OutletProduct, err error) {

	// get book by id
	if err = repo.DB.First(&data).Where("outlet_id = ? and product_id = ?", payload.OutletID, payload.ProductID).Error; err != nil {
		return
	}

	// update value
	data.Stock = payload.Stock
	data.Price = payload.Price
	data.IsActive = payload.IsActive

	// update book data
	err = repo.DB.Save(&data).Error
	return
}

func (repo *outletProductRepo) DeleteByID(OutletId uint, ProductId uint) (data dto.OutletProduct, err error) {
	// get book by id
	if err = repo.DB.First(&data).Where("outlet_id = ? and product_id = ?", OutletId, ProductId).Error; err != nil {
		return
	}

	// delete book
	err = repo.DB.Delete(&data).Error
	return
}
