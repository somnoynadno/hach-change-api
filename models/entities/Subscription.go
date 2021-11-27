package entities

import "hack-change-api/models/auxiliary"

type Subscription struct {
	auxiliary.BaseModel
	PublisherID  uint
	Publisher    User
	SubscriberID uint
	Subscriber   User
}
