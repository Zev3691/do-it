package middleware

import (
	"net/http"
	"re_new/util"
	"re_new/util/conf"
	"re_new/util/errorx"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var noAuthpath map[string]struct{}

func init() {
	noAuthpath = make(map[string]struct{}, 1)
	noAuthpath["/users/login"] = struct{}{}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		_, err := parseToken(auth)
		if err != nil {
			c.Abort()
			message := "token 过期" + err.Error()
			c.JSON(http.StatusOK, util.Response(errorx.ErrPermissionDeny, errorx.NewCustomErrMsg(message, "")))
		} else {
			println("token 正确")
		}
		c.Next()
	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.GetString("secret")), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
