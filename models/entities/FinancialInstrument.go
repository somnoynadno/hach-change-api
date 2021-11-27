package entities

import "hack-change-api/models/auxiliary"

type FinancialInstrument struct {
	auxiliary.BaseModel
	Ticker           string          `json:"ticker"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	InstrumentTypeID uint            `json:"instrumentTypeID"`
	InstrumentType   *InstrumentType `json:"instrumentType,omitempty"`
	BlogPosts        []*BlogPost     `json:"blogPosts,omitempty" gorm:"many2many:post_instruments;"`
}
