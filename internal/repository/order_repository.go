package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreateOrder(order *domain.Order) error {
	return database.DB.Create(order).Error
}

func CreateOrderItem(orderItem *domain.OrderItem) error {
	return database.DB.Create(orderItem).Error
}