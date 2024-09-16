package types

import (
	"time"

	"github.com/benosborntech/feedme/common/types"
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
	User         types.User  `json:"user"`
	ExpiresAt    time.Time   `json:"expiresAt"`
}
