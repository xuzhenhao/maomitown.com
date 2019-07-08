package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestID 为每一个请求加上requestID的中间件
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			u4, _ := uuid.NewV4()
			requestID = u4.String()
		}

		c.Set("X-Request-Id", requestID)
		//设置在返回包的Header中
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
