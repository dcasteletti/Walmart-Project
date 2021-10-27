package service

import (
	"fmt"
	"proyect/entity"
	repo "proyect/repository"
)

func GetAllProducts() (*entity.Products, error) {
	response, err := repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetProducts(value string) (*entity.Products, error) {
	response, err := repo.GetProducts(value)
	if err != nil {
		return nil, err
	}

	for _, r := range response.Product {
		fmt.Println(r.Id)
		fmt.Println(r.Brand)
		fmt.Println(r.Description)
		fmt.Println(r.Image)
		fmt.Println(r.Price)
		fmt.Println(".-...")
	}

	fmt.Println()

	return response, nil
}
