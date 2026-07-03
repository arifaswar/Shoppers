package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	// Product   Product   `gorm:"foreignKey:ProductID" json:"-"`
	Quantity int     `gorm:"not null" json:"quantity"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Total    float64 `gorm:"type:decimal(10,2);not null" json:"total"`
}

func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) error {
	orderItem.ID = uuid.New()
	return nil
}
