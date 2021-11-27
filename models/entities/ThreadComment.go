package entities

import "hack-change-api/models/auxiliary"

type ThreadComment struct {
	auxiliary.BaseModel
	Text      string
	CommentID uint
	Comment   *Comment `json:",omitempty"`
	AuthorID  uint
	Author    *User `json:",omitempty"`
}
