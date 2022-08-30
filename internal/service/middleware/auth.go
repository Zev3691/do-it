package middleware

import (
	"net/http"
	"re_new/repository/mysql"
	"re_new/util"
	"re_new/util/auth"
	"re_new/util/errorx"
	"re_new/util/log"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
			return
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		claim, err := parseToken(auth)
		if err != nil {
			c.Abort()
			message := "token 过期" + err.Error()
			c.JSON(http.StatusOK, util.Response(errorx.ErrPermissionDeny, errorx.NewCustomErrMsg(message, "")))
			return
		} else {
			log.Debugf(c, "claim %v", claim)
			log.Debugf(c, "claim.Id %v", claim.Id)
			if claim.Id == "" {
				c.Abort()
				c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
				return
			}
			uid, _ := strconv.Atoi(claim.Id)
			u, err := mysql.FindById(c, uid)
			if err != nil {
				log.Error(c, err.Error())
				c.Abort()
				c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
				return
			}
			if u == nil {
				log.Errorf(c, "无法找到用户 %v", claim.Id)
				c.Abort()
				c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
				return
			}
		}
		c.Next()
	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(auth.GetSecret()), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
