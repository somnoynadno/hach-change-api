package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email      string `gorm:"unique_index;not null;"`
	Password   string `json:"-" gorm:"not null;"`
	Name       string
	Surname    string
	IsVerified bool `gorm:"not null;default:false;"`
	LastLogin  *time.Time
	LastVisit  *time.Time
	BlogPosts  []*BlogPost
	Comments   []*Comment
	Followers  []*User `gorm:"many2many:user_followers"`
}
