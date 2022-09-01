package redis

import (
	"context"
	"encoding/json"
	"re_new/util/conf"
	"re_new/util/log"
	"time"
)

type Auth struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
}

func (value *Auth) SetAuthToken(ctx context.Context, key string) error {
	key = "auth " + key
	val, err := json.Marshal(*value)
	if err != nil {
		return err
	}
	exp := int(time.Second) * conf.GetInt("oneDayOfHours") * 2
	cmd := GetRedisDB().SetEX(ctx, key, val, time.Duration(exp))
	if cmd.Err() != nil {
		log.Error(ctx, cmd.Err().Error())
		return cmd.Err()
	}
	return nil
}

func (value *Auth) GetAuthToken(ctx context.Context, key string) (*Auth, error) {
	key = "auth " + key
	tmp, err := GetRedisDB().Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	if tmp == nil {
		if err := GetRedisDB().Del(ctx, key).Err(); err != nil {
			log.Error(ctx, err.Error())
			return nil, err
		}
	}
	if err := json.Unmarshal(tmp, value); err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return value, nil
}
