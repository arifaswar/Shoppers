package repository

import (
	"fmt"
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
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
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

func GetCartItem(cartID uuid.UUID, productID uuid.UUID) (*domain.CartItem, error) {

	var cartItem domain.CartItem

	fmt.Printf("Getting cart item for cartID: %s, productID: %s\n", cartID, productID)

	err := database.DB.
		Preload("Product").
		Preload("Cart").
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		First(&cartItem).
		Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func GetCartItemsById(itemID uuid.UUID) (*domain.CartItem, error) {
	var item domain.CartItem

	err := database.DB.
		Preload("Product").
		Preload("Cart").
		First(&item, "id = ?", itemID).
		Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateCartItem(item *domain.CartItem) error {
	return database.DB.Save(item).Error
}

func DeleteCartItem(itemID uuid.UUID) error {
	return database.DB.Delete(&domain.CartItem{}, "id = ?", itemID).Error
}

func GetCartItemWithCart(
	itemID string,
) (*domain.CartItem, error) {

	var item domain.CartItem

	err := database.DB.
		Preload("Cart").
		Preload("Product").
		First(
			&item,
			"id = ?",
			itemID,
		).
		Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}