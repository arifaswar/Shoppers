package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CartID uuid.UUID `gorm:"type:uuid;not null" json:"cart_id"`
	Cart Cart `gorm:"foreignKey:CartID" json:"-"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Product Product `gorm:"foreignKey:ProductID" json:"-"`
	Quantity int `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (cartItem *CartItem) BeforeCreate(tx *gorm.DB) error {
	cartItem.ID = uuid.New()
	return nil
}