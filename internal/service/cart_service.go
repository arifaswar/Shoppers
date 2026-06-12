package service

import (
	"errors"
	"fmt"
	"shoppers/internal/database"
	"shoppers/internal/domain"
	"shoppers/internal/dto"
	"shoppers/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddToCart(
	userID string, 
	productID string, 
	quantity int,
	) error {
		userUUID, err := uuid.Parse(userID)
		if err != nil {
			return err
		}
		productUUID, err := uuid.Parse(productID)
		if err != nil {
			return err
		}

		cart, err := repository.GetCartByUserId(userUUID.String())
		
		if err != nil {

			if !errors.Is(
				err,
				gorm.ErrRecordNotFound,
			) {
				return err
			}

			cart = &domain.Cart{
				UserID: userUUID,
			}

			if err := repository.CreateCart(cart); err != nil {
				return err
			}
		}

		existingItem, err := repository.GetCartItem(
			cart.ID, 
			productUUID,
		)

		fmt.Printf("Existing Item: %+v\n", err)
		
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

				newItem := domain.CartItem{
					CartID: cart.ID,
					ProductID: productUUID,
					Quantity: quantity,
				}
				product, err := repository.GetProductById(productID)
				if err != nil {
					return err
				}
				if quantity > product.Stock {
					return errors.New("quantity exceeds stock")
				}

				return repository.AddCartItem(&newItem)
			} 

			return err
		}

		existingItem.Quantity += quantity

		err = database.DB.Save(existingItem).Error
		fmt.Printf("Updated Item: %+v\n", err)
		if err != nil {
			return err
		}
		return nil
}

func GetCartItems(userID string) (*dto.CartResponse, error) {

	fmt.Printf("Getting cart items for userID: %s\n", userID)

	userUUID, err := uuid.Parse(userID)

	if err != nil {
		return nil, err
	}

	cart, err := repository.GetCartByUserId(userUUID.String())
	if err != nil {
		return nil, err
	}

	response := &dto.CartResponse{
		Items: []dto.CartItemResponse{},
	}
	totalPrice := 0.0

	fmt.Printf("Cart ID: %s\n", cart.ID)
	fmt.Printf("Items Count: %d\n", len(cart.Items))
	for _, item := range cart.Items {
		subtotal := float64(item.Quantity) * item.Product.Price
		response.Items = append(response.Items, dto.CartItemResponse{
			ID:          item.ID.String(),
			ProductID:   item.ProductID.String(),
			ProductName: item.Product.Name,
			Price:       item.Product.Price,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
	})
	totalPrice += subtotal
	}
	response.Total = totalPrice

	return response, nil
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

func UpdateCartItem(userID string, itemID string, quantity int) error {

	userUUID, err := uuid.Parse(userID)
	itemUUID, err := uuid.Parse(itemID)

	// item, err := repository.GetCartItemsById(itemID)
	if quantity < 1 {
		return errors.New("quantity must be greater than 0")
	}

	item, err := repository.GetCartItemWithCart(itemUUID.String())
	if err != nil {
		return err
	}

	if item.Cart.UserID != userUUID {
		return errors.New("unauthorized")
	}
	
	if quantity > item.Product.Stock {
		return errors.New("quantity exceeds stock")
	}

	item.Quantity = quantity

	return repository.UpdateCartItem(item)
}

func DeleteCartItem(userID string, itemID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	itemUUID, err := uuid.Parse(itemID)
	if err != nil {
		return err
	}

	item, err := repository.GetCartItemWithCart(itemUUID.String())
	if err != nil {
		return err
	}
	if item.Cart.UserID != userUUID {
		return errors.New("unauthorized")
	}

	return repository.DeleteCartItem(itemUUID)
}