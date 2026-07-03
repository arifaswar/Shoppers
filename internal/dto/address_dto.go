package dto

import "github.com/google/uuid"

type CreateAddressRequest struct {
	ID            uuid.UUID `gorm:"type:uuid"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	District      string    `gorm:"not null" json:"district"`
	City          string    `gorm:"not null" json:"city"`
	Province      string    `gorm:"not null" json:"province"`
	AddressLine1  string    `gorm:"not null" json:"address_line1"`
	AddressLine2  string    `gorm:"" json:"address_line2"`
	PostalCode    string    `gorm:"" json:"postal_code"`
	Country       string    `gorm:"" json:"country"`
	Label         string    `gorm:"" json:"label"`
	RecipientName string    `gorm:"" json:"recipient_name"`
	PhoneNumber   string    `gorm:"" json:"phone_number"`
	IsDefault     bool      `gorm:"" json:"is_default"`
}
