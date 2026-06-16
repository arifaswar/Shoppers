package service

import (
	"shoppers/internal/domain"
	"shoppers/internal/dto"
	"shoppers/internal/repository"

	"github.com/google/uuid"
)

func CreateAddress(userID string, req *dto.CreateAddressRequest) error {
	userUUID, err := uuid.Parse(userID)

	if err != nil {
		return err
	}

	address := domain.Address{
		UserID: userUUID,
		RecipientName: req.RecipientName,
		Label: req.Label,
		PhoneNumber: req.PhoneNumber,
		Country: req.Country,
		Province: req.Province,
		City: req.City,
		District: req.District,
		PostalCode: req.PostalCode,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		IsDefault: req.IsDefault,
	}

	return repository.CreateAddress(&address)
}

func GetAddressByUserID(userID string) ([]domain.Address, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return repository.GetAddressByUserID(userUUID.String())
}