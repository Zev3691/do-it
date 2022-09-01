package middleware

import (
	"context"
	"net/http"
	"re_new/repository/mysql"
	"re_new/repository/redis"
	"re_new/util"
	authx "re_new/util/auth"
	"re_new/util/errorx"
	"re_new/util/log"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 前端请求Header中设置的token保持不变，校验有效性以缓存中的token为准，千万不要直接校验Header中的token

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
			return
		}

		// 从缓存中获取，过期自动删除
		oldToken := getTokenVal(c, auth)
		if oldToken == nil {
			c.Abort()
			message := "token 过期, 请重新登录"
			c.JSON(http.StatusOK, util.Response(errorx.ErrPermissionDeny, errorx.NewCustomErrMsg(message, "")))
			return
		}

		// 校验token
		var newToken string
		claim, err := parseToken(strings.Fields(oldToken.Token)[1])
		if err != nil {
			// 内部token过期，重新生成
			if err.Errors == jwt.ValidationErrorExpired {
				var err error
				newToken, err = authx.NewToken(c, oldToken.UserName, 0, oldToken.UserId, "re_new", "relogin")
				if err != nil {
					c.Abort()
					message := "token 过期, " + err.Error()
					c.JSON(http.StatusOK, util.Response(errorx.ErrPermissionDeny, errorx.NewCustomErrMsg(message, "")))
					return
				}
			} else {
				c.Abort()
				message := "token 过期: " + err.Error()
				c.JSON(http.StatusOK, util.Response(errorx.ErrPermissionDeny, errorx.NewCustomErrMsg(message, "")))
				return
			}
		}

		// 新生成的token替换内部token，并重新获取一次校验值
		if newToken != "" {
			refTokenRedis(c, *oldToken, auth, newToken)
			claim, _ = parseToken(strings.Fields(newToken)[1])
		} else {
			// oldToken的值会随着newToken改变而改变，这里使用handler中的token作为固定key
			refTokenRedis(c, *oldToken, auth, oldToken.Token)
		}

		// 校验用户
		if claim.Id == "" {
			c.Abort()
			c.JSON(http.StatusOK, util.Response(errorx.ErrAccessFailed, errorx.NewCustomErrMsg("无法认证，重新登录", "")))
			return
		}
		uid, _ := strconv.Atoi(claim.Id)
		u, err1 := mysql.FindUserById(c, uid)
		if err1 != nil {
			log.Error(c, err1.Error())
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
		c.Next()
	}
}

func parseToken(token string) (*jwt.StandardClaims, *jwt.ValidationError) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(authx.GetSecret()), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err.(*jwt.ValidationError)
}

func getTokenVal(ctx context.Context, token string) *redis.Auth {
	value := redis.Auth{}
	val, err := value.GetAuthToken(ctx, token)
	if err != nil {
		return nil
	}
	return val
}

func refTokenRedis(ctx context.Context, val redis.Auth, oldToken, newToken string) {
	val.Token = newToken
	val.SetAuthToken(ctx, oldToken)
}
