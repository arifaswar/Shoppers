package dto

type UpdateCartItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required,min=1"`
}