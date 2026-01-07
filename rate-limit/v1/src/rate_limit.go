package ratelimit

import (
	"sync"
	"time"
)

// RateLimiter provides a simple rate limiter.
type RateLimiter struct {
	mu      sync.Mutex
	limit   int
	count   int
	resetAt time.Time
}

// New creates a new RateLimiter with the given limit per minute.
func New(limit int) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		resetAt: time.Now().Add(time.Minute),
	}
}

// Allow checks if the action is allowed under the rate limit.
func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.After(r.resetAt) {
		r.count = 0
		r.resetAt = now.Add(time.Minute)
	}

	if r.count < r.limit {
		r.count++
		return true
	}
	return false
}
