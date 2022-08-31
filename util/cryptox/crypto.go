package cryptox

import (
	"encoding/base64"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func GenSalt(length int) string {
	return RandStringBytesRmndr(length)
}

func Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
