package request

import (
	"github.com/google/uuid"
)

type Payment struct {
	TransactionId uuid.UUID `gorm:"type:uuid" json:"transaction_id"`
	Amount        float64   `json:"amount"`
	InsatllmentTo float32   `json:"installment_to"`
	Date          string    `json:"date"`
}
