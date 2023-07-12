package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLog() gin.HandlerFunc {
	// 日志中间件
	return func(c *gin.Context) {
		start := time.Now()
		log.Printf("request received %+v", c.Request.Body)
		c.Next()
		log.Printf("request processed, %s cost %+v", c.FullPath(), time.Since(start))
	}
}
