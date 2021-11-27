package entities

import "hack-change-api/models/auxiliary"

type InstrumentType struct {
	auxiliary.BaseModelCompact
	Name string
}
