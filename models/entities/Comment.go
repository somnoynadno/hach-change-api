package entities

import "hack-change-api/models/auxiliary"

type Comment struct {
	auxiliary.BaseModel
	Text           string           `json:"text"`
	BlogPostID     uint             `json:"blogPostID"`
	BlogPost       *BlogPost        `json:"blogPost,omitempty"`
	AuthorID       uint             `json:"authorID"`
	Author         *User            `json:"author,omitempty"`
	ThreadComments []*ThreadComment `json:"threadComments,omitempty"`
	Likes          []*User          `json:"likes,omitempty" gorm:"many2many:like_comments"`
}
