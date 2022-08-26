package user

import (
	"github.com/gin-gonic/gin"
)

func Registry(middlewares ...gin.HandlerFunc) func(engin *gin.Engine) {
	return func(engin *gin.Engine) {
		user := engin.Group("/users")
		user.Use(middlewares...)

		user.GET("/test", Test)
	}
}