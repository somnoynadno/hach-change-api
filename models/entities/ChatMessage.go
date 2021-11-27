package entities

import "hack-change-api/models/auxiliary"

type ChatMessage struct {
	auxiliary.BaseModel
	Text   string
	Seen   bool `gorm:"not null;default:false;"`
	FromID uint
	From   User
	ToID   uint
	To     User
}
