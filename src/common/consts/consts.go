package consts

import "time"

const (
	CHANNEL_PREFIX = "channel"
	LOCK_PREFIX    = "lock"
	POLLER_PREFIX  = "poller"

	LOCK_DURATION = 30 * time.Second

	WORLD_CIRCUMFERENCE = 40075000.0
	MIN_RADIUS          = 5     // metres
	MAX_RADIUS          = 20000 // metres
	MAX_PRECISION       = 16
)
