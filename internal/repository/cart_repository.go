package repository

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"

	"github.com/google/uuid"
)

func CreateCart(cart *domain.Cart) error {
	return database.DB.Create(cart).Error
}

func GetCartById(id string) (*domain.Cart, error) {
	var cart domain.Cart
	err := database.DB.
		Preload("Items").
		Where("id = ?", id).
		First(&cart).
		Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func GetCartByUserId(userId string) (*domain.Cart, error) {
	var cart domain.Cart

	err := database.DB.
		Preload("Items").
		Where("user_id = ?", userId).
		First(&cart).
		Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func AddCartItem(cartItem *domain.CartItem) error {
	return database.DB.Create(cartItem).Error
}

func GetCartWithItems(userID uuid.UUID) (*domain.Cart, error) {
	var cart domain.Cart
	err := database.DB.
		Preload("Items").
		Preload("Items.Product").
		Where("user_id = ?", userID).
		First(&cart).
		Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}