package models

import "github.com/google/uuid"

type Customer struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar(300)" json:"name"`
	Phone       float32   `json:"phone"`
	Address     string    `gorm:"type:text" json:"address"`
	Nik         string    `gorm:"type:text" json:"nik"`
}
