package dto

type OrderResponse struct {
	OrderID string `json:"order_id"`
	Total   float64 `json:"total"`
	Status  string `json:"status"`
}

type UpdateOrderStatusRequest struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}