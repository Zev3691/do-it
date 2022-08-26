package service

import (
	"re_new/internal/server/auth"
	"re_new/internal/server/user"
	"re_new/internal/service/middleware"

	"github.com/gin-gonic/gin"
)

var engin *gin.Engine

func HTTP() {
	engin = gin.Default()
	engin.Use(middleware.Cors())

	auth.Registry()(engin)
	user.Registry(middleware.Auth())(engin)

	if err := engin.Run(":9090"); err != nil {
		panic(err)
	}
}
