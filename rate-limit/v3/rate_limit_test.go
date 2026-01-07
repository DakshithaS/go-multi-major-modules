package ratelimit_test

import (
	"testing"
	"time"

	ratelimit "github.com/DakshithaS/go-multi-major-modules/rate-limit/v3/src"
)

func TestRateLimiter(t *testing.T) {
	rl := ratelimit.New(2, time.Minute)
	if !rl.Allow("key1") {
		t.Error("First allow for key1 should succeed")
	}
	if !rl.Allow("key1") {
		t.Error("Second allow for key1 should succeed")
	}
	if rl.Allow("key1") {
		t.Error("Third allow for key1 should fail")
	}

	// Test with different window
	rl2 := ratelimit.New(1, time.Second)
	if !rl2.Allow("key") {
		t.Error("Allow should succeed")
	}
	time.Sleep(1100 * time.Millisecond)
	if !rl2.Allow("key") {
		t.Error("After window, allow should succeed")
	}
}
