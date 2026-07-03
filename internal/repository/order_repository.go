package repository

import (
	"fmt"
	"shoppers/internal/database"
	"shoppers/internal/domain"
)

func CreateOrder(order *domain.Order) error {
	return database.DB.Create(order).Error
}

func CreateOrderItem(orderItem *domain.OrderItem) error {
	return database.DB.Create(orderItem).Error
}

func GetAllOrders() ([]domain.Order, error) {

	fmt.Println("Getting all orders")
	var orders []domain.Order
	err := database.DB.
		Preload("OrderItems").
		// Preload("OrderItems.Product").
		// Preload("User").
		Order("created_at DESC").
		Find(&orders).
		Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrdersByUserID(userID string) ([]domain.Order, error) {

	// fmt.Printf("Getting orders for user %s", userID)
	var orders []domain.Order

	err := database.DB.
		Preload("OrderItems").
		// Preload("OrderItems.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).
		Error

	fmt.Printf("Orders for user %s: %v", userID, err)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderByID(orderID string) (*domain.Order, error) {
	fmt.Printf("Getting order by ID: %s", orderID)
	var order domain.Order
	err := database.DB.
		Preload("OrderItems").
		// Preload("OrderItems.Product").
		Where("id = ?", orderID).
		First(&order).
		Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func UpdateOrder(order *domain.Order) error {
	return database.DB.Save(order).Error
}

func UpdateOrderStatus(orderID string, status domain.OrderStatus) error {
	return database.DB.Model(&domain.Order{}).
		Where("id = ?", orderID).
		Update("status", status).
		Error
}

func DeleteOrder(orderID string) error {
	return database.DB.Where("id = ?", orderID).Delete(&domain.Order{}).Error
}
