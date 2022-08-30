package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var defaultSecret = "3f6e5ff13309da890021f4a97b8780be23581e3ffc9df558edb0837fad9e4aa8bf39d55fa12b147ea3053cf9ca602481b492764c61eb72cea0781ba57fd8a6f2"

func GetSecret() string {
	return defaultSecret
}

func NewClaims(audience string, exp int64, id, issuer, subject string) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Audience:  audience,          // 受众
		ExpiresAt: exp,               // 失效时间
		Id:        id,                // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    issuer,            // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   subject,           // 主题
	}
}
