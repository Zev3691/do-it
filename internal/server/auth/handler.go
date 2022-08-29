package auth

import (
	"re_new/internal/server"
	"re_new/repository/mysql"
	"re_new/util/conf"
	"re_new/util/errorsx"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.Response(c, errorsx.ErrJsonParse, err.Error())
		return
	}
	// 获取用户
	user := &mysql.User{}
	user.Name = req.Name
	user.Password = req.Password
	err := user.LoginMath(c)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			server.Response(c, errorsx.ErrAccessFailed, "用户不存在")
		} else {
			server.Response(c, errorsx.ErrAccessFailed, err.Error())
		}
		return
	}
	if req.Password == user.Password {
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
			server.Success(c, "登录成功", "Bearer "+token)
			return
		} else {
			server.Response(c, errorsx.ErrAccessFailed, err.Error())
			return
		}
	} else {
		server.Response(c, errorsx.ErrAccessFailed, "登录失败")
		return
	}
}
