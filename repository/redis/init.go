package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisDb *redis.Client

func Init() {
	redisDb = initRedis()
}

func Close() {
	redisDb.Close()
}

func initRedis() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0, // use default DB
	})
	if cmd := db.Ping(context.Background()); cmd.Err() != nil {
		panic(fmt.Sprintf("redis init fail, err: %v", cmd.Err()))
	}
	return db
}
