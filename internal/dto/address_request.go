package dto

type CreateAddressRequest struct {
	UserID  string `json:"user_id"`
	District string `json:"district"`
	City    string `json:"city"`
	Province string `json:"province"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	PostalCode string `json:"postal_code"`
	Country string `json:"country"`
	Label   string `json:"label"`
	RecipientName string `json:"recipient_name"`
	PhoneNumber string `json:"phone_number"`
	IsDefault   bool   `json:"is_default"`
}