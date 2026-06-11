package service

import (
	"shoppers/internal/domain"
	"shoppers/internal/repository"
)

func CreateCategory(name string) error {
	category := domain.Category{
		Name: name,
	}

	return repository.CreateCategory(&category)
}

func GetCategories() ([]domain.Category, error) {
	return repository.GetCategories()
}

func GetCategoryById(id string) (*domain.Category, error) {
	return repository.GetCategoryById(id)
}