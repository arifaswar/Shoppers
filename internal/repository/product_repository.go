package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreateProduct(product *domain.Product) error {
	return database.DB.Create(product).Error
}

func GetProducts() ([]domain.Product, error) {
	var products []domain.Product

	err := database.DB.Find(&products).Error

	if err !=nil {
		return nil, err
	}

	return products, nil
}