package service

import (
	"Routes/src/banco"
	"Routes/src/models"
	"Routes/src/repository"
	"errors"
)

func CreateProductService(productName string, costPrice float64, sellPrice float64) (*models.Product, error) {
	product := &models.Product{
		ProductName:  productName,
		CostPrice:    costPrice,
		SellPrice:    sellPrice,
		Quantity:     0,
		QuantitySell: 0,
	}

	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	productRepository := repository.NewProductRepository(db)

	product, err = productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	productRepository := repository.NewProductRepository(db)

	return productRepository.GetAllProducts()
}

func FindProductById(productId int64) (*models.Product, error) {
	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	productRepository := repository.NewProductRepository(db)

	product, err := productRepository.GetProductById(productId)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProductById(productId int64, productVO models.ProductRequestVO) (*models.Product, error) {
	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	productRepository := repository.NewProductRepository(db)

	product := &models.ProductRequestVO{
		ProductName: productVO.ProductName,
		CostPrice:   productVO.CostPrice,
		SellPrice:   productVO.SellPrice,
	}

	productUpdated, err := productRepository.UpdateProductParsingId(productId, product)
	if err != nil {
		return nil, err
	}

	return productUpdated, nil
}

func DeleteProductParsingId(productId int64) error {
	db, err := banco.Conectar()
	if err != nil {
		return errors.New("Erro ao conectar com o banco!")
	}

	productRepository := repository.NewProductRepository(db)
	if err := productRepository.DeleteProductById(productId); err != nil {
		return err
	}

	return nil
}
