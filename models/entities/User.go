package entities

import (
	"hack-change-api/models/auxiliary"
	"time"
)

type User struct {
	auxiliary.BaseModel
	Email       string      `json:"email" gorm:"unique_index;not null;"`
	Password    string      `json:"-" gorm:"not null;"`
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Username    string      `json:"username" gorm:"unique_index;not null;"`
	IsVerified  bool        `json:"isVerified" gorm:"not null;default:false;"`
	LastLogin   *time.Time  `json:"lastLogin"`
	LastVisit   *time.Time  `json:"lastVisit"`
	BlogPosts   []*BlogPost `json:"blogPosts,omitempty"`
	Comments    []*Comment  `json:"comments,omitempty"`
	Publishers  []*User     `gorm:"many2many:subscriptions" json:"publishers,omitempty"`
	Subscribers []*User     `gorm:"many2many:subscriptions" json:"subscribers,omitempty"`
}
