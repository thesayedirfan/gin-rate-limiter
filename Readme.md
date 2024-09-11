# Gin Rate Limiter with Token Bucket and Fixed Window Algorithms

**⚠️ This project is experimental and should not be used in production.**

This project demonstrates a rate-limiting middleware for the [Gin](https://gin-gonic.com/) web framework in Go. It includes two rate-limiting algorithms: Token Bucket and Fixed Window. The middleware is easily extensible and can be integrated with various rate-limiting strategies.

## Features

- **Token Bucket Rate Limiter**: Limits the number of requests within a specific time frame, refilling tokens at a fixed rate.
- **Fixed Window Rate Limiter**: Limits the number of requests in a fixed time window.

## Prerequisites

- [Go 1.19+](https://golang.org/dl/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

# Rate Limiting Algorithms 

### Token Bucket Rate Limiter
The TokenBucketRateLimiter allows requests as long as tokens are available in the bucket. The bucket refills at a fixed rate over time. If the bucket is empty, requests are denied until more tokens are available.

Example usage in main.go:

```go

tokenBucketLimiter := algorithms.NewTokenBucketRateLimiter(5, 1 * time.Second)
r.Use(middleware.RateLimitingMiddlerware(tokenBucketLimiter, "too many requests"))

```
This example allows 5 requests per second using the token bucket strategy.


### Fixed Window Rate Limiter

The FixedWindowRateLimiter limits requests by grouping them into fixed-length time windows. Each window has a maximum number of requests allowed. If the limit is exceeded, subsequent requests in that window are rejected.

Example usage in main.go:

```go
fixedWindowLimiter := algorithms.NewFixedWindowRateLimiter(5, 1 * time.Second)
r.Use(middleware.RateLimitingMiddlerware(fixedWindowLimiter, "too many requests"))
```
This example allows 5 requests per second within each fixed window.