package auth

import (
	"context"
	"re_new/repository/mysql"
	"re_new/util/conf"
	"re_new/util/errorx"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Name     string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Biz struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *Biz {
	return &Biz{
		Db: db,
	}
}

func (srv *Biz) Login(ctx context.Context, req *LoginRequest) (string, error) {
	// 获取用户
	user := &mysql.User{}
	user.Name = req.Name
	user.Password = req.Password
	err := user.LoginMath(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg("用户不存在"))
		} else {
			return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg(err.Error()))
		}
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
			// success return
			return "Bearer " + token, nil
		} else {
			return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg(err.Error()))
		}
	} else {
		return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg("登录失败"))
	}
}
