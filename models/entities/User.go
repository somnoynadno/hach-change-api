package entities

import (
	"hack-change-api/models/auxiliary"
	"time"
)

type User struct {
	auxiliary.BaseModel
	Email       string `gorm:"unique_index;not null;"`
	Password    string `json:"-" gorm:"not null;"`
	Name        string
	Surname     string
	IsVerified  bool `gorm:"not null;default:false;"`
	LastLogin   *time.Time
	LastVisit   *time.Time
	BlogPosts   []*BlogPost
	Comments    []*Comment
	Publishers  []*User `gorm:"many2many:subscriptions"`
	Subscribers []*User `gorm:"many2many:subscriptions"`
}
