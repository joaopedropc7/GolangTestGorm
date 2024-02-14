package repository

import (
	"Routes/src/models"
	"errors"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := r.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(productId int64) (models.Product, error) {
	var product models.Product
	if err := r.DB.First(&product, productId).Error; err != nil {
		erro := "Não foi encontrado nenhum registro com este ID!"
		return models.Product{}, errors.New(erro)
	}
	return product, nil
}

func (r *ProductRepository) UpdateProductParsingId(productId int64, productUpdated *models.ProductRequestVO) (*models.Product, error) {

	var product models.Product
	if err := r.DB.First(&product, productId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Não foi encontrado nenhum registro com este ID!")
		}
		return nil, err
	}

	product.ProductName = productUpdated.ProductName
	productUpdated.CostPrice = productUpdated.CostPrice
	productUpdated.SellPrice = productUpdated.SellPrice

	if err := r.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) DeleteProductById(productId int64) error {

	var existingProduct models.Product
	err := r.DB.First(&existingProduct, productId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Produto não encontrado para exclusão")
		}
		return err
	}

	return r.DB.Delete(&existingProduct).Error
}
