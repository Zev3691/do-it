package auth

import (
	"net/http"
	"re_new/repository/mysql"
	"re_new/util"
	"re_new/util/conf"
	"re_new/util/errorsx"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username     string
	Userpassword string
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, util.Response(errorsx.ErrJsonParse))
		return
	}

	// 获取用户
	user := &mysql.User{}
	err := user.LoginMath(c)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, util.Response(errorsx.ErrAccessFailed, errorsx.NewCustomErrMsg("用户不存在", "")))
		} else {
			c.JSON(http.StatusOK, util.Response(errorsx.ErrAccessFailed, errorsx.NewCustomErrMsg(err.Error(), "")))
		}
		return
	}
	if req.Userpassword == user.Password {
		expiresTime := time.Now().Unix() + int64(conf.GetInt("oneDayOfHours"))
		claims := jwt.StandardClaims{
			Audience:  user.Name,             // 受众
			ExpiresAt: expiresTime,           // 失效时间
			Id:        strconv.Itoa(user.ID), // 编号
			IssuedAt:  time.Now().Unix(),     // 签发时间
			Issuer:    "gin hello",           // 签发人
			NotBefore: time.Now().Unix(),     // 生效时间
			Subject:   "login",               // 主题
		}
		var jwtSecret = []byte(conf.GetString("secret"))
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
			c.JSON(http.StatusOK, util.Success("登录成功", "Bearer "+token))
			return
		} else {
			c.JSON(http.StatusOK, util.Response(errorsx.ErrAccessFailed, errorsx.NewCustomErrMsg(err.Error(), "")))
			return
		}
	} else {
		c.JSON(http.StatusOK, util.Response(errorsx.ErrAccessFailed, errorsx.NewCustomErrMsg("登录失败", "")))
		return
	}
}
