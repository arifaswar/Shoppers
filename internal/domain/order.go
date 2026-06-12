package domain
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	Total     float64   `gorm:"type:decimal(10,2)" json:"total"`
	Status    string    `gorm:"type:varchar(50)" json:"status"`
}

type OrderStatusPending struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:char(36)" json:"order_id"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	Total     float64   `gorm:"type:decimal(10,2)" json:"total"`
	Status    string    `gorm:"type:varchar(50)" json:"status"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New()
	return nil
}