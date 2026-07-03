package dto

import (
	"github.com/google/uuid"
	)

type CheckoutRequest struct {
	AddressID uuid.UUID `gorm:"type:uuid;not null" json:"address_id"`
}

type CheckoutResponse struct {
	OrderID uuid.UUID `gorm:"type:uuid;primaryKey" json:"order_id"`
	Total float64 `json:"total"`
	Status string `json:"status"`
}