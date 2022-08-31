package user

import (
	"github.com/gin-gonic/gin"
)

func Registry(middlewares ...gin.HandlerFunc) func(engin *gin.Engine) {
	return func(engin *gin.Engine) {
		user := engin.Group("/users")
		user.Use(middlewares...)

		user.POST("/create", Create)
		user.POST("/update", Update)
		user.POST("/list", List)
		user.POST("/delete", Delete)
		user.POST("/get", Get)
		user.POST("/pwd", GetUserPwd)
	}
}
