package types

import "time"

type Business struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	CreatedBy   int       `json:"createdBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedAt   time.Time `json:"createdAt"`
}
