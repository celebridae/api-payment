package service

import (
	"api-payment/internal/database"
	"api-payment/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducties() ([]*entity.Product, error) {

	products, err := ps.ProductDB.GetProducties()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {

	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (ps *ProductService) PostProduct(prod entity.Product) (string, error) {
	product, err := ps.ProductDB.PostProduct(prod)
	if err != nil {
		return "", err
	}
	return product, nil

}
