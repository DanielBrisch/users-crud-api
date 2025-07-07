package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		userID, exists := c.Get("user_id")

		entry := log.WithFields(logrus.Fields{
			"status":   status,
			"method":   method,
			"path":     path,
			"ip":       clientIP,
			"duration": duration.Milliseconds(),
		})

		if exists {
			entry = entry.WithField("user_id", userID)
		}

		if status >= 500 {
			entry.Error("HTTP request")
		} else if status >= 400 {
			entry.Warn("HTTP request")
		} else {
			entry.Info("HTTP request")
		}
	}
}
