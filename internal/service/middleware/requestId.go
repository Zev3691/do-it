package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.NewV4().String()
		c.Set("request_id", requestId)
		c.Writer.Header().Set("request_id", requestId)

		// ctxLog := log.WithValues("request_id", requestId)
		// c.Set(log.GetContextKey(), ctxLog)

		c.Next()
	}
}
