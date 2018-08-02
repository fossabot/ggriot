package ggriot

import (
	"time"
)

var (
	// CSRs is the current requested calls in the short term rate limit.
	CSRs int64
	// CLRs is the current requested calls in the long term rate limit.
	CLRs int64
	// StopRateLimit can be called to disable checking rate limit.
	StopRateLimit = make(chan struct{})
	// RateLimitingActive is a bool to see if there's already a limiter active
	RateLimitingActive bool
)

func init() {
	RateLimitingActive = true
	go ShortRateTime()
	go LongRateTime()
}

// ShortRateTime is the timer that loops every second to check.
func ShortRateTime() {
	for {
		select {
		default:
			time.Sleep(2050 * time.Millisecond)
			CSRs = 0
		case <-StopRateLimit:
			CSRs = 0
			RateLimitingActive = false
			return
		}
	}
}

// LongRateTime is the timer that loops every 2 minutes to check/clear rate .
func LongRateTime() {
	for {
		select {
		default:
			time.Sleep(121000 * time.Millisecond)
			CLRs = 0
		case <-StopRateLimit:
			CLRs = 0
			RateLimitingActive = false
			return
		}
	}
}

// CheckRateLimit check to see if we have reached threshold.
func CheckRateLimit() bool {
	if RateLimitingActive == false {
		return true
	}
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
