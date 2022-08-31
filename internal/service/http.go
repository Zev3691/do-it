package service

import (
	"re_new/internal/server/auth"
	"re_new/internal/server/user"
	"re_new/internal/service/middleware"
	"re_new/util/log"

	"github.com/gin-gonic/gin"
)

var engin *gin.Engine

func HTTP() {
	engin = gin.New()
	engin.Use(middleware.RequestId(), middleware.Cors())
	engin.Use(middleware.Data(), middleware.Log(), middleware.Recovery(log.GetOriginLog(), true))

	auth.Registry()(engin)
	user.Registry(middleware.Auth())(engin)

	if err := engin.Run(":9090"); err != nil {
		panic(err)
	}
}
