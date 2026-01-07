package ratelimit_test

import (
	"testing"
	"time"

	"github.com/dakshithas/go-multi-major-modules/rate-limit/v2/src"
)

func TestRateLimiter(t *testing.T) {
	rl := ratelimit.New(2)
	if !rl.Allow("key1") {
		t.Error("First allow for key1 should succeed")
	}
	if !rl.Allow("key1") {
		t.Error("Second allow for key1 should succeed")
	}
	if rl.Allow("key1") {
		t.Error("Third allow for key1 should fail")
	}
	if !rl.Allow("key2") {
		t.Error("First allow for key2 should succeed")
	}

	// Wait for reset
	time.Sleep(time.Minute + time.Second)
	if !rl.Allow("key1") {
		t.Error("After reset, allow for key1 should succeed")
	}
}
