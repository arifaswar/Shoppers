package dto

type CartItemResponse struct {
	ID string `json:"id"`
	ProductID string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}

type CartResponse struct {
	// ID string `json:"id"`
	// UserID string `json:"user_id"`
	Items []CartItemResponse `json:"items"`
	Total float64 `json:"total"`
}