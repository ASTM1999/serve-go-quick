package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Logger เป็น middleware ที่ใช้ในการบันทึก log ของ request และ response
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		// บันทึก log ของ request และ response ได้ตามต้องการ
		// ตัวอย่างเช่น fmt.Printf("Request %s %s %s\n", c.Request.Method, c.Request.URL.String(), latency)
	}
}
