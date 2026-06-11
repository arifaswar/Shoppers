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

func GetProductById(id string) (*domain.Product, error) {
	return repository.GetProductById(id)
}

func UpdateProduct(
	id string,
	name string,
	description string,
	price float64,
	stock int,
	imageURL string,
) error {

	product, err := repository.GetProductById(id)
	if err != nil {
		return err
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock
	product.ImageURL = imageURL

	return repository.UpdateProduct(product)
}

func DeleteProduct(id string) error {
	return repository.DeleteProduct(id)
}