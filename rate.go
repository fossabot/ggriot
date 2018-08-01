package ggriot

import (
	"time"
)

var (
	// CSRs is the current requested calls in the short term rate limit.
	CSRs int64
	// CLRs is the current requested calls in the long term rate limit.
	CLRs int64
)

func init() {
	go ShortRateTime()
	go LongRateTime()
}

// ShortRateTime is the timer that loops every second to check.
func ShortRateTime() {
	for {
		time.Sleep(2050 * time.Millisecond)
		CSRs = 0
	}
}

// LongRateTime is the timer that loops every 2 minutes to check/clear rate .
func LongRateTime() {
	for {
		time.Sleep(121000 * time.Millisecond)
		CLRs = 0
	}
}

// CheckRateLimit check to see if we have reached threshold.
func CheckRateLimit() bool {
	if CSRs >= ShortLimit {
		return false
	}
	if CLRs >= LongLimit {
		return false
	}
	CSRs++
	CLRs++
	return true
}
