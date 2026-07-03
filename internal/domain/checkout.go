package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkout struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID  `gorm:"type:uuid;not null"`
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"not null"`
}

func (checkout *Checkout) BeforeCreate(tx *gorm.DB) error {
	checkout.ID = uuid.New()
	return nil
}

