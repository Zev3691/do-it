package redis

import (
	"context"
	"fmt"
	"re_new/util/cryptox"
	"time"
)

func Auth(ctx context.Context, userId int, userName, token string) bool {
	key := cryptox.UserName(userId, userName)
	return redisDb.Get(ctx, key).String() == token
}

func AddUser(ctx context.Context, userCook, token string) error {
	msg, err := redisDb.HSet(ctx, "auth", userCook, token, time.Duration(60*60)).Result()
	defer fmt.Println(msg)
	if err != nil {
		return err
	}
	return nil
}
