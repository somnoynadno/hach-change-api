package entities

import "hack-change-api/models/auxiliary"

type ChatMessage struct {
	auxiliary.BaseModel
	Text   string `json:"text"`
	Seen   bool   `json:"seen" gorm:"not null;default:false;"`
	FromID uint   `json:"fromID"`
	From   *User  `json:"from,omitempty"`
	ToID   uint   `json:"toID"`
	To     *User  `json:"to,omitempty"`
}
