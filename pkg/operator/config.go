package operator

import (
	"time"
)

// Config is the operator configuraton.
type Config struct {
	// Resync is the time the controller will resync its resources.
	ResyncDuration time.Duration
}
