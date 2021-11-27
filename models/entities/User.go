package entities

import (
	"hack-change-api/models/auxiliary"
	"time"
)

type User struct {
	auxiliary.BaseModel
	Email        string         `json:"email" gorm:"unique_index;not null;"`
	Password     string         `json:"-" gorm:"not null;"`
	Name         string         `json:"name"`
	Surname      string         `json:"surname"`
	Username     string         `json:"username" gorm:"unique_index;not null;"`
	About        *string        `json:"about"`
	IsVerified   bool           `json:"isVerified" gorm:"not null;default:false;"`
	LastLogin    *time.Time     `json:"lastLogin"`
	LastVisit    *time.Time     `json:"lastVisit"`
	BlogPosts    []*BlogPost    `json:"blogPosts,omitempty" gorm:"foreignKey:AuthorID"`
	Comments     []*Comment     `json:"comments,omitempty" gorm:"foreignKey:AuthorID"`
	ChatMessages []*ChatMessage `json:"chatMessages,omitempty" gorm:"foreignKey:FromID"`
	Publishers   []*User        `json:"publishers,omitempty" gorm:"many2many:subscriptions;ForeignKey:id;References:PublisherID;"`
	Subscribers  []*User        `json:"subscribers,omitempty" gorm:"many2many:subscriptions;ForeignKey:id;References:SubscriberID;"`
}
