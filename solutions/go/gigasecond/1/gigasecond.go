package gigasecond

import (
	"time"
)

// AddGigasecond adds 10^9 seconds to the given Time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
