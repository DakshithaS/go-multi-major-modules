package ratelimit_test

import (
	"testing"
	"time"

	ratelimit "github.com/DakshithaS/go-multi-major-modules/rate-limit/src"
)

func TestRateLimiter(t *testing.T) {
	rl := ratelimit.New(2)
	if !rl.Allow() {
		t.Error("First allow should succeed")
	}
	if !rl.Allow() {
		t.Error("Second allow should succeed")
	}
	if rl.Allow() {
		t.Error("Third allow should fail")
	}

	// Wait for reset
	time.Sleep(time.Minute + time.Second)
	if !rl.Allow() {
		t.Error("After reset, allow should succeed")
	}
}
