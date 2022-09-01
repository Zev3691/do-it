package repository

import (
	"re_new/repository/mongo"
	"re_new/repository/mysql"
	"re_new/repository/redis"
)

func Init() {
	mongo.Init()
	mysql.Init()
	redis.Init()
}
