package request

import (
	"github.com/google/uuid"
)

type CreteProduct struct {
	ProductName string    `gorm:"type:varchar(300)" json:"product_name"`
	Type        string    `gorm:"type:text" json:"type"`
	Description string    `gorm:"type:text" json:"description"`
	CategoryId  uuid.UUID `gorm:"type:uuid" json:"categoryId"`
}
