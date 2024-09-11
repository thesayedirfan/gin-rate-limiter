package algorithms_test



import (
	"testing"
	"time"

	"github.com/thesayedirfan/gin-rate-limiter/pkg/algorithms"
)


func TestTokenBucket(t *testing.T){


	const TIME_IN_SECONDS = 1 * time.Second

	rateLimiter := algorithms.NewTokenBucketRateLimiter(1,TIME_IN_SECONDS)

	ip := "192.168.0.1"

	if rateLimiter.Allow(ip) == false {
		t.Errorf("failed")
	}

	if rateLimiter.Allow(ip) == true {
		t.Errorf("failed")
	}

	time.Sleep(TIME_IN_SECONDS)

	if rateLimiter.Allow(ip) == false {
		t.Errorf("failed")
	}

}