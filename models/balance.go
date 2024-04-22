package models

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Number      int       `gorm:"autoIncrement;uniqueIndex:idx_number" json:"number"`
	SourceId    uuid.UUID `gorm:"type:uuid"`
	CashIn      int       `gorm:"type:int" json:"cash_in"`
	CashOut     int       `gorm:"type:int" json:"cash_out"`
	Saldo       int       `gorm:"type:int" json:"saldo"`
	Description string    `gorm:"type:varchar(300)" json:"description"`
	Date        time.Time `json:"date"`
	Source      Source    `gorm:"foreignKey:SourceId"`
}
