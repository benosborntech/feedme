package types

import (
	"time"
)

type Token struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	TokenType    ServiceType `json:"tokenType"`
	UserId       int         `json:"userId"`
	ExpiresAt    time.Time   `json:"expiresAt"`
}
