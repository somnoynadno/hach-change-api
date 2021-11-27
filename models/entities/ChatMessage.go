package entities

import "hack-change-api/models/auxiliary"

type ChatMessage struct {
	auxiliary.BaseModel
	Text   string
	Seen   bool `gorm:"not null;default:false;"`
	FromID uint
	From   *User `json:",omitempty"`
	ToID   uint
	To     *User `json:",omitempty"`
}
