package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreateUser(user *domain.User) error {
	return database.DB.Create(user).Error
}

func FindUserByEmail(email string) (*domain.User, error) {

	var user domain.User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, err
}