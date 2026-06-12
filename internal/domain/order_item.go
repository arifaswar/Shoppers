package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type OrderItem struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:char(36)" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:char(36)" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Total     float64   `gorm:"type:decimal(10,2);not null" json:"total"`
}

func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) error {
	orderItem.ID = uuid.New()
	return nil
}