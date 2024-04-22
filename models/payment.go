package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id            uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	TransactionId uuid.UUID   `gorm:"type:uuid"`
	CustomerId    uuid.UUID   `gorm:"type:uuid"`
	ProductId     uuid.UUID   `gorm:"type:uuid"`
	Amount        float64     `json:"amount"`
	InsatllmentTo float32     `json:"installment_to"`
	Date          time.Time   `json:"date"`
	Transaction   Transaction `gorm:"foreignKey:TransactionId"`
	Customer      Customer    `gorm:"foreignKey:CustomerId"`
	Product       Product     `gorm:"foreignKey:ProductId"`
}
