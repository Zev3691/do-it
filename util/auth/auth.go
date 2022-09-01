package auth

import (
	"context"
	"re_new/repository/redis"
	"re_new/util/conf"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var defaultSecret = "3f6e5ff13309da890021f4a97b8780be23581e3ffc9df558edb0837fad9e4aa8bf39d55fa12b147ea3053cf9ca602481b492764c61eb72cea0781ba57fd8a6f2"

func GetSecret() string {
	return defaultSecret
}

func NewToken(ctx context.Context, audience string, exp int64, id int, issuer, subject string) (string, error) {
	expr := exp
	if expr == 0 {
		expr = time.Now().Unix() + int64(conf.GetInt("oneDayOfHours"))
	}
	claims := &jwt.StandardClaims{
		Audience:  audience,          // 受众
		ExpiresAt: expr,              // 失效时间
		Id:        strconv.Itoa(id),  // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    issuer,            // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   subject,           // 主题
	}

	jwtSecret := []byte(GetSecret())
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := tokenClaims.SignedString(jwtSecret)
	if err == nil {
		token := "Bearer " + t
		return token, nil
	}
	return "", err
}

func SetToRedis(ctx context.Context, token, userName string, userId int) error {
	value := redis.Auth{
		Token:    token,
		UserName: userName,
		UserId:   userId,
	}
	return value.SetAuthToken(ctx, token)
}
