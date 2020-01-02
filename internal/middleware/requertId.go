package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request_id", "123a")
	}
}
