package dto

import (
	"github.com/google/uuid"
)

type UpdateOrderStatusRequest struct {
	OrderID uuid.UUID `gorm:"type:uuid;primaryKey" json:"order_id"`
	Status  string `json:"status"`
}

type OrderResponse struct {
	OrderID uuid.UUID `gorm:"type:uuid;primaryKey" json:"order_id"`
	Total   float64 `json:"total"`
	Status  string `json:"status"`
}