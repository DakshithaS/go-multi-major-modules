# rate-limit/v2

This is version 2 of the rate-limit module.

## API

- `New(limit int) *RateLimiter`: Creates a new rate limiter with the given limit per minute per key.
- `Allow(key string) bool`: Checks if the action is allowed for the given key.

## Breaking Changes

- Changed `Allow() bool` to `Allow(key string) bool` from v1, now supports per-key limiting.