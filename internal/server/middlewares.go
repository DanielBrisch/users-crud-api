package server

import (
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func MiddlewareLogger() gin.HandlerFunc {
	return gin.Logger()
}

func MiddlewareRecovery() gin.HandlerFunc {
	return gin.Recovery()
}

func MiddlewareCORS() gin.HandlerFunc {
	return cors.Default()
}

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5)
		visitors[ip] = limiter
	}
	return limiter
}

func MiddlewareRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getVisitor(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	}
}
