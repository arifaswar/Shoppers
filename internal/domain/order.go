package domain
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User      User `gorm:"foreignKey:UserID" json:"-"`
	Total     float64   `gorm:"type:decimal(10,2)" json:"total"`
	Status    OrderStatus `gorm:"type:varchar(20)" json:"status"`
	RecipientName string `gorm:"type:varchar(255)" json:"recipient_name"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	Country   string `gorm:"type:varchar(100)" json:"country"`	
	Province  string `gorm:"type:varchar(100)" json:"province"`
	City      string `gorm:"type:varchar(100)" json:"city"`
	District  string `gorm:"type:varchar(100)" json:"district"`
	PostalCode   string `gorm:"type:varchar(20)" json:"postal_code"`
	AddressLine1 string `gorm:"type:varchar(255)" json:"address_line1"`
	AddressLine2 string `gorm:"type:varchar(255)" json:"address_line2"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	Payment Payment `gorm:"foreignKey:OrderID" json:"payment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderStatus string
	const (
	OrderPending    OrderStatus = "pending"
	OrderPaid       OrderStatus = "paid"
	OrderProcessing OrderStatus = "processing"
	OrderShipped    OrderStatus = "shipped"
	OrderCompleted  OrderStatus = "completed"
	OrderCancelled  OrderStatus = "cancelled"
)

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New()
	return nil
}