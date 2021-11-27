package entities

import "hack-change-api/models/auxiliary"

type BlogPost struct {
	auxiliary.BaseModel
	Text        string
	Instruments []*FinancialInstrument `gorm:"many2many:post_instruments;"`
	AuthorID    uint
	Author      User
	Comments    []*Comment
}
