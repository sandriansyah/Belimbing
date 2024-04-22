package models

import (
	"github.com/google/uuid"
)

type Source struct {
	Id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(300)" json:"name"`
}
