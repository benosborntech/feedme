package types

import "time"

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
