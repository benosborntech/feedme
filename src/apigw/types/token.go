package types

import (
	"time"
)

type UserInfo struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Token struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	TokenType    ServiceType `json:"tokenType"`
	UserId       int         `json:"user"`
	ExpiresAt    time.Time   `json:"expiresAt"`
}
