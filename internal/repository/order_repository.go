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

func GetOrdersByUserID(userID string) ([]domain.Order, error) {
	var orders []domain.Order

	err := database.DB.
		Preload("items").
		Preload("items.product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).
		Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderByID(orderID string) (*domain.Order, error) {
	var order domain.Order
	err := database.DB.
		Preload("items").
		Preload("items.product").
		Preload("user").
		Where("id = ?", orderID).
		First(&order).
		Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func UpdateOrderStatus(orderID string, status domain.OrderStatus) error {
	return database.DB.Model(&domain.Order{}).
		Where("id = ?", orderID).
		Update("status", status).
		Error
}