package auth

import (
	"re_new/internal/biz/auth"
	"re_new/internal/server"
	"re_new/repository/mysql"
	"re_new/util/errorx"
	"re_new/util/validata"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}

	if err := validata.Struct(req); err != nil {
		server.Response(c, errorx.New(errorx.ErrInvalidParameter, errorx.NewMsg(err.Error())))
		return
	}

	srv := auth.New(mysql.NewMysqlDB(c))
	token, err := srv.Login(c, &req)
	if err != nil {
		server.Response(c, err)
		return
	}
	server.Success(c, "", token)
}
