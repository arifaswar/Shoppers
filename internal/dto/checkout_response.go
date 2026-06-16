package dto

type CheckoutResponse struct {
	OrderID string `json:"order_id"`
	Total float64 `json:"total"`
	Status string `json:"status"`
}