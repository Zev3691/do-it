package user

import (
	"re_new/util/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Test(c *gin.Context) {
	log.Debug("debug", zap.Any("ddddd", "versionnnn"))
	c.JSON(200, "ok")
}
