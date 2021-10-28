package service

import (
	"proyect/entity"
	"proyect/repository"
)

func GetAllProducts() (*entity.Products, error) {
	response, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetProducts(value string, isID bool) (*entity.Products, error) {
	var err error
	products := &entity.Products{}

	if isID {
		products, err = repository.GetProductByID(value)
	} else {
		products, err = repository.GetProducts(value)
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}