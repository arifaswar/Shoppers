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

func GetProductById(id string) (*domain.Product, error) {
	var product domain.Product

	err := database.DB.
		Where("id = ?", id).
		First(&product).
		Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProduct(product *domain.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id string) error {
	return database.DB.Delete(&domain.Product{}, "id = ?", id).Error
}