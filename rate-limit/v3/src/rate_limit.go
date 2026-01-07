package ratelimit

import (
	"sync"
	"time"
)

// RateLimiter provides a per-key rate limiter with configurable window.
type RateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	clients map[string]*clientLimiter
}

type clientLimiter struct {
	count   int
	resetAt time.Time
}

// New creates a new RateLimiter with the given limit and window.
func New(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]*clientLimiter),
	}
}

// Allow checks if the action is allowed for the given key.
func (r *RateLimiter) Allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	cl, exists := r.clients[key]
	if !exists || now.After(cl.resetAt) {
		cl = &clientLimiter{resetAt: now.Add(r.window)}
		r.clients[key] = cl
	}

	if cl.count < r.limit {
		cl.count++
		return true
	}
	return false
}
