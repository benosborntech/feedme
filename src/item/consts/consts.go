package consts

import "time"

const (
	PORT = "3010"

	POLL_FREQ = 5 * time.Second

	DB_HISTORY_WINDOW = 12 * time.Hour
)
