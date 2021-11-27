package entities

import "hack-change-api/models/auxiliary"

type Subscription struct {
	auxiliary.BaseModel
	PublisherID  uint  `json:"publisherID"`
	Publisher    *User `json:"publisher,omitempty"`
	SubscriberID uint  `json:"subscriberID"`
	Subscriber   *User `json:"subscriber,omitempty"`
}
