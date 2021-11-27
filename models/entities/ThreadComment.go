package entities

import "hack-change-api/models/auxiliary"

type ThreadComment struct {
	auxiliary.BaseModel
	Text      string   `json:"text"`
	CommentID uint     `json:"commentID"`
	Comment   *Comment `json:"comment,omitempty"`
	AuthorID  uint     `json:"authorID"`
	Author    *User    `json:"author,omitempty"`
	Likes     []*User  `json:"likes,omitempty" gorm:"many2many:like_thread_comments"`
}
