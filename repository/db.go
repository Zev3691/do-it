package repository

import (
	"re_new/repository/mongo"
	"re_new/repository/mysql"
)

func Init() {
	mongo.Init()
	mysql.Init()
	// redis.Init()
}
