# rate-limit/v1

This is version 1 of the rate-limit module.

## API

- `New(limit int) *RateLimiter`: Creates a new rate limiter with the given limit per minute.
- `Allow() bool`: Checks if the action is allowed.

## Breaking Changes

This is the initial version, no breaking changes.