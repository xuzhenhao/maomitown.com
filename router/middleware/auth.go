package middleware

import (
	"github.com/gin-gonic/gin"
	"maomitown.com/handler"
	"maomitown.com/pkg/errno"
	"maomitown.com/pkg/token"
)

// AuthMiddleware token鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.TokenInvalidError, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
