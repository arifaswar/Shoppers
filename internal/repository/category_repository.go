package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreateCategory(category *domain.Category) error {
	return database.DB.Create(category).Error
}

func GetCategories() ([]domain.Category, error) {
	var categories []domain.Category

	err := database.DB.
		Find(&categories).
		Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryById(id string) (*domain.Category, error) {
	var category domain.Category

	err := database.DB.
		Where("id = ?", id).
		First(&category).
		Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}