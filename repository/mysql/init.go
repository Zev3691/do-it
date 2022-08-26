package mysql

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = initMysql()
}

func initMysql() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/re_new?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("初始化数据库失败： %v", err))
	}
	return db
}

func NewMysqlDB(ctx context.Context) *gorm.DB {
	// TODO 从ctx中获取logger实例并赋值给session中的logger
	return db.Session(&gorm.Session{
		Context: ctx,
	})
}
