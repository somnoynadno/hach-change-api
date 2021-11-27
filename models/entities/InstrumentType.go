package entities

import "hack-change-api/models/auxiliary"

type InstrumentType struct {
	auxiliary.BaseModelCompact
	Name        string                 `json:"name"`
	Instruments []*FinancialInstrument `json:"instruments,omitempty"`
}
