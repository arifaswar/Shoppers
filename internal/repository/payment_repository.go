package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreatePayment(payment *domain.Payment) error {
	return database.DB.Create(payment).Error
}

func GetPaymentByOrderID(orderID string) (*domain.Payment, error) {
	var payment domain.Payment

	err:= database.DB.
		Where("order_id = ?", orderID).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func UpdatePayment(payment *domain.Payment) error {
	return database.DB.Save(payment).Error
}