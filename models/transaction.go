package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CustomerId       uuid.UUID `gorm:"type:uuid"`
	ProductId        uuid.UUID `gorm:"type:uuid"`
	SourceId         uuid.UUID `gorm:"type:uuid"`
	DownPayment      int       `gorm:"type:int" json:"down_payment"`
	InstallmentValue int       `gorm:"type:int" json:"installment_value"`
	InstallmentCount int       `gorm:"type:int" json:"installment_count"`
	Date             time.Time `json:"date"`
	PurchasePrice    int       `json:"purchase_price"`
	Customer         Customer  `gorm:"foreignKey:CustomerId"`
	Product          Product   `gorm:"foreignKey:ProductId"`
	Source           Source    `gorm:"foreignKey:SourceId"`
}
