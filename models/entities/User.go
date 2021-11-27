package entities

import (
	"hack-change-api/models/auxiliary"
	"time"
)

type User struct {
	auxiliary.BaseModel
	Email         string `gorm:"unique_index;not null;"`
	Password      string `json:"-" gorm:"not null;"`
	Name          string
	Surname       string
	IsVerified    bool `gorm:"not null;default:false;"`
	LastLogin     *time.Time
	LastVisit     *time.Time
	BlogPosts     []*BlogPost
	Comments      []*Comment
	Followers     []*User `gorm:"many2many:followers_subscriptions;ForeignKey:follower_id;References:follower_id"`
	Subscriptions []*User `gorm:"many2many:followers_subscriptions;ForeignKey:subscription_id;References:subscription_id"`
}
