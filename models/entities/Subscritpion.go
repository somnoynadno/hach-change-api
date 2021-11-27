package entities

import "hack-change-api/models/auxiliary"

type Subscription struct {
	auxiliary.BaseModel
	FollowerID     uint
	Follower       User
	SubscriptionID uint
	Subscription   User
}
