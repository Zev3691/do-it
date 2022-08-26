package main

import (
	"re_new/internal/service"
	"re_new/repository"
	"re_new/util"
	"re_new/util/conf"
	"re_new/util/log"
	"re_new/util/validata"

	"go.uber.org/zap"
)

func main() {
	// 资源初始化
	conf.Init()
	log.Init()
	repository.Init()
	validata.Init()

	// 调试信息打印
	debugPrint()

	// http || rpc 服务启动
	service.Init()
}

func debugPrint() {
	// test和dev环境执行的逻辑
	switch util.GetVersion() {
	case util.Test, util.Development:
		log.Info("配置文件加载key: %v", zap.Any("viper,key: ", conf.GetAllKey()))
	}
}
