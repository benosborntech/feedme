package types

import "time"

type Item struct {
	Id        int       `json:"id"`
	Location  string    `json:"location"`
	ItemType  int       `json:"itemType"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}
