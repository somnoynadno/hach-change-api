package entities

import "hack-change-api/models/auxiliary"

type Comment struct {
	auxiliary.BaseModel
	Text           string
	BlogPostID     uint
	BlogPost       *BlogPost `json:",omitempty"`
	AuthorID       uint
	Author         *User            `json:",omitempty"`
	ThreadComments []*ThreadComment `json:",omitempty"`
}
