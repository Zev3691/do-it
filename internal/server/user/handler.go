package user

import (
	"re_new/internal/biz/user"
	"re_new/internal/server"
	"re_new/util/errorx"
	"re_new/util/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Test(c *gin.Context) {
	log.Debug("debug", zap.Any("ddddd", "versionnnn"))
	c.JSON(200, "ok")
}

func Create(c *gin.Context) {
	var req user.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}
	if err := user.Create(c, &req); err != nil {
		server.Response(c, err)
		return
	}
	server.SuccessNIL(c)
}
