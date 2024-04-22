package models

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	ProductName string           `gorm:"type:varchar(300)" json:"product_name"`
	Type        string           `gorm:"type:text" json:"type"`
	Description string           `gorm:"type:text" json:"description"`
	CategoryId  uuid.UUID        `gorm:"type:uuid"`
	Category    Product_category `gorm:"foreignKey:CategoryId"`
}
