package types

import "time"

type Item struct {
	Id         int       `json:"id"`
	Location   string    `json:"location"`
	ItemType   int       `json:"itemType"`
	Quantity   int       `json:"quantity"`
	ExpiresAt  time.Time `json:"expiresAt"`
	CreatedBy  int       `json:"createdBy"`
	BusinessId int       `json:"businessId"`
	UpdatedAt  time.Time `json:"updatedAt"`
	CreatedAt  time.Time `json:"createdAt"`
}
