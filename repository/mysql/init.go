package mysql

import (
	"context"
	"fmt"
	"re_new/util/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = initMysql()
}

func initMysql() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/re_new?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.NewGormLogger(),
	})
	if err != nil {
		panic(fmt.Sprintf("初始化数据库失败： %v", err))
	}
	println("[mysql] Successfully connected.")
	return db
}

func NewMysqlDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
