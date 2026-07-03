package service

import (
	"shoppers/internal/domain"
	"shoppers/internal/dto"
	"shoppers/internal/repository"

	"github.com/google/uuid"
)

func CreatePayment(req dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {
	order, err := repository.GetOrderByID(req.OrderID)
	if err != nil {
		return nil, err
	}
	
	payment := domain.Payment{
		OrderID: order.ID,
		Method:  domain.PaymentMethod(req.Method),
		Status:  domain.PaymentStatusPending,
		Amount:  req.Amount,
		TransactionID: uuid.New().String(),
	}

	err = repository.CreatePayment(&payment)
	if err != nil {
		return nil, err
	}

	return &dto.PaymentResponse{
		ID:              payment.ID,
		OrderID:         payment.OrderID,
		Method:          string(payment.Method),
		Status:          string(payment.Status),
		Amount:          payment.Amount,
		TransactionID: payment.TransactionID,
	}, nil
}