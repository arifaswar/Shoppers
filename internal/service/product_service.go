package service

import (
	"shoppers/internal/domain"
	"shoppers/internal/repository"
)

func CreateProduct(
	name string,
	description string,
	price float64,
	stock int,
	imageURL string,
) error {
	product := domain.Product{
		Name: name,
		Description: description,
		Price: price,
		Stock: stock,
		ImageURL: imageURL,
	}

	return repository.CreateProduct(&product)
}

func GetProducts() ([]domain.Product, error) {
	return repository.GetProducts()
}