package user

import (
	"re_new/internal/biz/user"
	"re_new/internal/server"
	"re_new/util/errorx"

	"github.com/gin-gonic/gin"
)

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

func Update(c *gin.Context) {
	var req user.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}
	if err := user.Update(c, &req); err != nil {
		server.Response(c, err)
		return
	}
	server.SuccessNIL(c)
}

func List(c *gin.Context) {
	var req user.ListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}
	item, err := user.List(c, &req)
	if err != nil {
		server.Response(c, err)
		return
	}

	server.Success(c, "", item)
}

func Delete(c *gin.Context) {
	var req user.DeleteRequset
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}

	if err := user.Delete(c, &req); err != nil {
		server.Response(c, err)
		return
	}

	server.SuccessNIL(c)
}

func Get(c *gin.Context) {
	var req user.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}

	data, err := user.Get(c, &req)
	if err != nil {
		server.Response(c, err)
		return
	}

	server.Success(c, "", data)
}

func GetUserPwd(c *gin.Context) {
	var req user.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.CustomResponse(c, errorx.ErrJsonParse, err.Error())
		return
	}
	pwd, err := user.GetUserPwd(c, &req)
	if err != nil {
		server.Response(c, err)
		return
	}
	server.Success(c, "", pwd)
}
