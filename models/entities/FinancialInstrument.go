package entities

import "hack-change-api/models/auxiliary"

type FinancialInstrument struct {
	auxiliary.BaseModel
	Ticker           string
	Name             string
	Description      string
	InstrumentTypeID uint
	InstrumentType   InstrumentType
	Posts            []*FinancialInstrument `gorm:"many2many:post_instruments;"`
}
