package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)



type RateLimitingStatergy interface {
	Allow(ip string) bool
}

func RateLimitingMiddlerware(strategy RateLimitingStatergy,msg string) gin.HandlerFunc {
	return func (c *gin.Context)  {
		ip := c.ClientIP()
		if !strategy.Allow(ip){
			c.JSON(http.StatusTooManyRequests,gin.H{"error":msg})
			c.Abort()
			return
		}
		c.Next()
	}
}
