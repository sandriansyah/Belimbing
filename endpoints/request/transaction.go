package request

import (
	"time"

	"github.com/google/uuid"
)

type CreateTransaction struct {
	Id               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CustomerId       uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductId        uuid.UUID `gorm:"type:uuid;primaryKey"`
	SourceId         uuid.UUID `gorm:"type:uuid;primaryKey"`
	DownPayment      int       `json:"down_payment"`
	InstallmentValue int       `json:"installment_value"`
	InstallmentCount int       `json:"installment_count"`
	Date             time.Time `json:"date"`
	PurchasePrice    int       `json:"purchase_price"`
}
