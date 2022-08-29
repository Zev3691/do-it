package server

import (
	"net/http"
	"re_new/util"
	"re_new/util/errorsx"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, errCode errorsx.ErrCode, errMsg ...string) {
	var msg errorsx.CustomErrMsg
	if len(errMsg) == 1 {
		msg = errorsx.NewCustomErrMsg(errMsg[0], "")
	} else {
		msg = errorsx.NewCustomErrMsg(errMsg[0], errMsg[1])
	}
	c.JSON(http.StatusOK, util.Response(errCode, msg))
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, util.Success(msg, data))
}

func SuccessNIL(c *gin.Context) {
	c.JSON(http.StatusOK, util.SuccessNIL())
}
