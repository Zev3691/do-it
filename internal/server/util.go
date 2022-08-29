package server

import (
	"net/http"
	"re_new/util"
	"re_new/util/errorx"

	"github.com/gin-gonic/gin"
)

func CustomResponse(c *gin.Context, errCode errorx.ErrCode, errMsg ...string) {
	var msg errorx.CustomErrMsg
	if len(errMsg) == 1 {
		msg = errorx.NewCustomErrMsg(errMsg[0], "")
	} else {
		msg = errorx.NewCustomErrMsg(errMsg[0], errMsg[1])
	}
	c.JSON(http.StatusOK, util.Response(errCode, msg))
}

func Response(c *gin.Context, err error) {
	c.JSON(http.StatusOK, err)
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, util.Success(msg, data))
}

func SuccessNIL(c *gin.Context) {
	c.JSON(http.StatusOK, util.SuccessNIL())
}
