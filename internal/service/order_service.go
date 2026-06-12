package service

import (
	"errors"
	"shoppers/internal/database"
	"shoppers/internal/dto"
	"shoppers/internal/repository"
	"shoppers/internal/domain"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

func Checkout(userID string) (*dto.CheckoutResponse, error) {

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	cart, err := repository.GetCartByUserId(userUUID.String())
	if err != nil {
		return nil, err
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	var response *dto.CheckoutResponse
	err = database.DB.Transaction(func(tx *gorm.DB) error {

		order := domain.Order{
			ID:     uuid.New(),
			UserID: userUUID,
			Status: ("pending"),
		}

		total := 0.0
		for _, item := range cart.Items {
			subTotal := float64(item.Quantity) * item.Product.Price
			total += subTotal
		}

		order.Total = total

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for _, item := range cart.Items {

			subTotal := float64(item.Quantity) * item.Product.Price
			orderItem := domain.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Price:     item.Product.Price,
				Quantity:  item.Quantity,
				Total:     subTotal,
			}

			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}
		if err := tx.
			Where("cart_id = ?", cart.ID).
			Delete(&domain.CartItem{}).
			Error; err != nil {

			return err
		}

		response = &dto.CheckoutResponse{
			OrderID: order.ID.String(),
			Total:   order.Total,
			Status:  string(order.Status),
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}