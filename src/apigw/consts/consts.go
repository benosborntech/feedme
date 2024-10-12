package consts

import "time"

const (
	PORT = "3000"

	STATE_LENGTH   = 32
	SESSION_LENGTH = 32
	SESSION_EXPIRY = 5 * time.Minute
	JWT_ISSUER     = "feedme"
)
