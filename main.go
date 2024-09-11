package main

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/gin-rate-limiter/pkg/algorithms"
	"github.com/thesayedirfan/gin-rate-limiter/pkg/middleware"
)

func main() {

	r := gin.Default()

	// Example usage of TokenBucketRateLimiter
	tokenBucketLimiter := algorithms.NewTokenBucketRateLimiter(5,1 * time.Second)
	r.Use(middleware.RateLimitingMiddlerware(tokenBucketLimiter, "too many request"))


	fixedWindowLimiter := algorithms.NewFixedWindowRateLimiter(5,1 * time.Second)

	r.Use(middleware.RateLimitingMiddlerware(fixedWindowLimiter,"too many request"))


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
