package dto
import (
	"github.com/google/uuid"
)
type CreatePaymentRequest struct {
	OrderID string `json:"order_id" binding:"required"`
	Method  string `json:"method" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

type PaymentResponse struct {
	ID uuid.UUID `json:"id"`
	OrderID uuid.UUID `json:"order_id"`
	Method string `json:"method"`
	Status string `json:"status"`
	Amount float64 `json:"amount"`
	TransactionID string `json:"transaction_id"`
	PaidAt *string `json:"paid_at,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}