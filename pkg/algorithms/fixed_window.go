package algorithms

import (
	"sync"
	"time"

	"github.com/thesayedirfan/gin-rate-limiter/pkg/middleware"
)

type FixedWindow struct {
	mu              sync.Mutex
	lastRequestTime map[string]time.Time
	requestCount    map[string]int
	count           int
	duration        time.Duration
}

func NewFixedWindowRateLimiter(count int, duration time.Duration) middleware.RateLimitingStatergy {
	return &FixedWindow{
		lastRequestTime: make(map[string]time.Time),
		requestCount:    make(map[string]int),
		count:           count,
		duration:        duration,
	}
}

func (f *FixedWindow) Allow(ip string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now()
	
	lastTime, ok := f.lastRequestTime[ip]

	if !ok || now.Sub(lastTime) >= f.duration {
		f.lastRequestTime[ip] = now
		f.requestCount[ip] = 1
		return true
	}

	if f.requestCount[ip] < f.count {
		f.requestCount[ip]++
		return true
	}
	return false
}
