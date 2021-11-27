package entities

import "hack-change-api/models/auxiliary"

type Comment struct {
	auxiliary.BaseModel
	Text       string
	BlogPostID uint
	BlogPost   BlogPost
	AuthorID   uint
	Author     User
}
