package algorithms

import (
	"sync"
	"time"

	"github.com/thesayedirfan/gin-rate-limiter/pkg/middleware"
)

type tokenBucketLimiter struct {
	mu              sync.Mutex
	lastRequestTime map[string]time.Time
	tokenCount      map[string]int
	duration        time.Duration
	refillRate      int
}

// Constructor
func NewTokenBucketRateLimiter(refillRate int, duration time.Duration) middleware.RateLimitingStatergy {
	return &tokenBucketLimiter{
		lastRequestTime: make(map[string]time.Time),
		tokenCount:      make(map[string]int),
		refillRate:      refillRate,
		duration:        duration,
	}
}

func (t *tokenBucketLimiter) Allow(ip string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now()

	lastTime, ok := t.lastRequestTime[ip]
	if !ok {
		t.lastRequestTime[ip] = now
		t.tokenCount[ip] = t.refillRate - 1
		return true
	}

	elapsed := now.Sub(lastTime)

	tokensToAdd := int(elapsed / t.duration) * t.refillRate
	if tokensToAdd > 0 {
		t.tokenCount[ip] += tokensToAdd
		if t.tokenCount[ip] > t.refillRate {
			t.tokenCount[ip] = t.refillRate
		}
		t.lastRequestTime[ip] = now
	}

	if t.tokenCount[ip] > 0 {
		t.tokenCount[ip]--
		return true
	}
	
	return false
}
