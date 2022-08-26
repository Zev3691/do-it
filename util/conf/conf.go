package conf

import (
	"fmt"
	"re_new/util"

	"github.com/spf13/viper"
)

var conf *viper.Viper

func Init() {
	switch util.GetVersion() {
	case util.Development:
		viper.AddConfigPath("./conf")
		viper.SetConfigName("./dev.yml")
	case util.Production:
		viper.AddConfigPath("./conf")
		viper.SetConfigName("./production.yml")
	default:
		viper.AddConfigPath("../conf")
		viper.SetConfigType("yml")
		viper.SetConfigName("test")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("初始化配置文件失败： %v \n", err))
	}
	conf = viper.GetViper()
}

func GetString(key string) string {
	return conf.GetString(key)
}

func GetStringOrDefault(key, def string) string {
	v := conf.GetString(key)
	if v == "" {
		return def
	}
	return v
}

func GetBool(key string) bool {
	return conf.GetBool(key)
}

func GetStringSlice(key string) []string {
	return conf.GetStringSlice(key)
}

func GetInt(key string) int {
	return conf.GetInt(key)
}

func GetIntOrDefault(key string, def int) int {
	v := conf.GetInt(key)
	if v == 0 {
		return def
	}
	return v
}

func GetIntSlice(key string) []int {
	return conf.GetIntSlice(key)
}

func GetStringMap(key string) map[string]interface{} {
	return conf.GetStringMap(key)
}

func GetAllKey() []string {
	return conf.AllKeys()
}
