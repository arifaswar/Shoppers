package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreateAddress(address *domain.Address) error {
	return database.DB.Create(address).Error
}

func GetAddressByUserID(userID string) ([]domain.Address, error) {

	var address []domain.Address

	err := database.DB.
		Where("user_id = ?", userID).
		Find(&address).
		Error
	if err != nil {
		return nil, err
	}

	return address, nil
}

func GetAddressByID(id string) (*domain.Address, error) {

	var address domain.Address
	err := database.DB.
		Where("id = ?", id).
		First(&address).
		Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}