package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed PaymentStatus = "failed"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodPayPal PaymentMethod = "paypal"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

type Payment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	Order *Order `gorm:"foreignKey:OrderID" json:"order"`
	Method PaymentMethod `gorm:"type:varchar(50)"`
	Status PaymentStatus `gorm:"type:varchar(50)" json:"status"`
	Amount float64 `gorm:"type:decimal(10,2)" json:"amount"`
	TransactionID string `gorm:"type:varchar(100)" json:"transaction_id"`
	PaidAt *time.Time `gorm:"type:timestamp" json:"paid_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}