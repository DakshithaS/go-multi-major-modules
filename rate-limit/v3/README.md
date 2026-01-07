# rate-limit/v3

This is version 3 of the rate-limit module.

## API

- `New(limit int, window time.Duration) *RateLimiter`: Creates a new rate limiter with the given limit and time window.
- `Allow(key string) bool`: Checks if the action is allowed for the given key.

## Breaking Changes

- Changed `New(limit int)` to `New(limit int, window time.Duration)` from v2, now accepts configurable window.