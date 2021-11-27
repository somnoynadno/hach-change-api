package entities

import "hack-change-api/models/auxiliary"

type LikeThreadComment struct {
	auxiliary.BaseModelCompact
	ThreadCommentID uint           `json:"threadCommentID"`
	ThreadComment   *ThreadComment `json:"threadComment,omitempty"`
	UserID          uint           `json:"userID"`
	User            *User          `json:"user,omitempty"`
}
