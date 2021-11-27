package entities

import "hack-change-api/models/auxiliary"

type ThreadComment struct {
	auxiliary.BaseModel
	Text      string
	CommentID uint
	Comment   Comment
	AuthorID  uint
	Author    User
}
