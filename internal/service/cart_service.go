package service

import (
	"shoppers/internal/domain"
	"shoppers/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shoppers/internal/dto"
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
			if err == gorm.ErrRecordNotFound {
				newCart := domain.Cart{
					UserID: userUUID,
				}
				err = repository.CreateCart(&newCart)
				if err != nil {
					return err
				}
				cart = &newCart
			} else {
				return err
			}
		}
		item := domain.CartItem{
			CartID: cart.ID,
			ProductID: productUUID,
			Quantity: quantity,
		}
		return repository.AddCartItem(&item)
}

func GetCart(userID string) (*dto.CartResponse, error) {

	userUUID, err := uuid.Parse(userID)

	if err != nil {
		return nil, err
	}

	cart, err := repository.GetCartWithItems(userUUID)
	if err != nil {
		return nil, err
	}

	response := &dto.CartResponse{}
	totalPrice := 0.0

	for _, item := range cart.Items {
		subtotal := float64(item.Quantity) * item.Product.Price
		response.Items = append(
			response.Items, dto.CartItemResponse{
			ID: item.ID.String(),
			ProductID: item.ProductID.String(),
			ProductName: item.Product.Name,
			Price: item.Product.Price,
			Quantity: item.Quantity,
			Subtotal: subtotal,
		})
		totalPrice += subtotal
	}
	response.Total = totalPrice

	return response, nil
}