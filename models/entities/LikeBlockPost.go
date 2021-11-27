package entities

import "hack-change-api/models/auxiliary"

type LikeBlogPost struct {
	auxiliary.BaseModelCompact
	BlogPostID uint      `json:"blogPostID"`
	BlogPost   *BlogPost `json:"blogPost,omitempty"`
	UserID     uint      `json:"userID"`
	User       *User     `json:"user,omitempty"`
}
