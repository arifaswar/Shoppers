package dto

import "github.com/google/uuid"

type UpdateCartItemRequest struct {
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required,min=1"`
}

type CartItemResponse struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}

type CartResponse struct {
	// ID string `json:"id"`
	// UserID string `json:"user_id"`
	Items []CartItemResponse `json:"items"`
	Total float64            `json:"total"`
}