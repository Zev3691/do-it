package auth

import (
	"context"
	"re_new/repository/mysql"
	"re_new/util/auth"
	"re_new/util/errorx"
	"re_new/util/log"

	"go.uber.org/zap"
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
	log.Debug(ctx, "req ", zap.Any("req ", req))
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
		token, err := auth.NewToken(ctx, user.Name, 0, user.ID, "re_new", "login")
		if err != nil {
			return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg(err.Error()))
		}
		if err := auth.SetToRedis(ctx, token, user.Name, user.ID); err != nil {
			return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg(err.Error()))
		}
		return token, nil
	} else {
		return "", errorx.New(errorx.ErrAccessFailed, errorx.NewMsg("登录失败, 检查密码"))
	}
}
