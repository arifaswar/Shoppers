package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	User	  User `gorm:"foreignKey:UserID" json:"-"` 
	Label     string `gorm:"type:varchar(100)" json:"label"`
	RecipientName string `gorm:"type:varchar(255)" json:"recipient_name"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	Country   string `gorm:"type:varchar(100)" json:"country"`
	Province  string `gorm:"type:varchar(100)" json:"province"`
	City      string `gorm:"type:varchar(100)" json:"city"`
	District  string `gorm:"type:varchar(100)" json:"district"`
	PostalCode   string `gorm:"type:varchar(20)" json:"postal_code"`
	AddressLine1 string `gorm:"type:varchar(255)" json:"address_line1"`
	AddressLine2 string `gorm:"type:varchar(255)" json:"address_line2"`
	IsDefault bool   `gorm:"default:false" json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func(address *Address) BeforeCreate(tx *gorm.DB) error {
	address.ID = uuid.New()
	return nil
}