package service

import (
	"errors"
	"shoppers/internal/database"
	"shoppers/internal/domain"
	"shoppers/internal/dto"
	"shoppers/internal/repository"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func Checkout(userID string, addressID string) (*dto.CheckoutResponse, error) {

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	address, err := repository.GetAddressByID(addressID)
	if err != nil {
		return nil, err
	}

	if address.UserID != userUUID {
		return nil, errors.New("address does not belong to user")
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
		total := 0.0
		for _, item := range cart.Items {
			subTotal := float64(item.Quantity) * item.Product.Price
			total += subTotal
		}

		order := domain.Order{
			ID:            uuid.New(),
			UserID:        userUUID,
			Status:        domain.OrderStatus(domain.OrderPending),
			RecipientName: address.RecipientName,
			PhoneNumber:   address.PhoneNumber,
			Country:       address.Country,
			Province:      address.Province,
			City:          address.City,
			District:      address.District,
			PostalCode:    address.PostalCode,
			AddressLine1:  address.AddressLine1,
			AddressLine2:  address.AddressLine2,
			Total:         total,
		}

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

		if err := tx.Where("cart_id = ?", cart.ID).Delete(&domain.CartItem{}).Error; err != nil {
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

func GetMyOrders(userID string) ([]dto.OrderResponse, error) {

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	orders, err := repository.GetOrdersByUserID(userUUID.String())
	if err != nil {
		return nil, err
	}
	var response []dto.OrderResponse
	for _, order := range orders {
		response = append(response, dto.OrderResponse{
			OrderID: order.ID.String(),
			Total:   order.Total,
			Status:  string(order.Status),
		})
	}
	return response, nil
}

func GetOrderByID(orderID string) (*dto.OrderResponse, error) {

	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, err
	}
	order, err := repository.GetOrderByID(orderUUID.String())
	if err != nil {
		return nil, err
	}
	response := &dto.OrderResponse{
		OrderID: order.ID.String(),
		Total:   order.Total,
		Status:  string(order.Status),
	}
	return response, nil
}

func UpdateOrderStatus(orderID string, status string) error {

	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	order.Status = domain.OrderStatus(status)

	return repository.UpdateOrderStatus(order.ID.String(), order.Status)
}
