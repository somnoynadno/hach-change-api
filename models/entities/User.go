package entities

import (
	"hack-change-api/models/auxiliary"
	"time"
)

type User struct {
	auxiliary.BaseModel
	Email       string `gorm:"unique_index;not null;"`
	Password    string `gorm:"not null;" json:"-"`
	Name        string
	Surname     string
	IsVerified  bool `gorm:"not null;default:false;"`
	LastLogin   *time.Time
	LastVisit   *time.Time
	BlogPosts   []*BlogPost `json:",omitempty"`
	Comments    []*Comment  `json:",omitempty"`
	Publishers  []*User     `gorm:"many2many:subscriptions" json:",omitempty"`
	Subscribers []*User     `gorm:"many2many:subscriptions" json:",omitempty"`
}
