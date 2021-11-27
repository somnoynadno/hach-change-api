package entities

import "hack-change-api/models/auxiliary"

type LikeComment struct {
	auxiliary.BaseModelCompact
	CommentID uint     `json:"commentID"`
	Comment   *Comment `json:"comment,omitempty"`
	UserID    uint     `json:"userID"`
	User      *User    `json:"user,omitempty"`
}
