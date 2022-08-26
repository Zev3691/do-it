package auth

import (
	"github.com/gin-gonic/gin"
)

func Registry(middlewares ...gin.HandlerFunc) func(engin *gin.Engine) {
	return func(engin *gin.Engine) {
		user := engin.Group("/auth")
		user.Use(middlewares...)

		user.POST("/login", Login)
	}
}
