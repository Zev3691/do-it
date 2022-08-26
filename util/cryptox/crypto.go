package cryptox

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Token(name, password, salt string) string {
	tims := strconv.Itoa(int(time.Now().Unix()))
	str := name + password + salt + tims
	strSlc := strings.Split(str, "")
	sort.Strings(strSlc)
	str = strings.Join(strSlc, "")

	h := sha1.New()
	h.Write([]byte(str))
	tokenByte := h.Sum(nil)

	return fmt.Sprintf("%x", tokenByte)
}

func UserName(userId int, userName string) string {
	str := strconv.Itoa(userId) + userName
	strSlc := strings.Split(str, "")
	sort.Strings(strSlc)
	str = strings.Join(strSlc, "")

	h := sha1.New()
	h.Write([]byte(str))
	cryptoUserName := h.Sum(nil)

	return fmt.Sprintf("%x", cryptoUserName)
}
