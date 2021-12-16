package main

import (
	"xirang/model"
	"xirang/public"
	"xirang/public/tools"
	"xirang/router"

	"github.com/spf13/viper"
)

// Init 初始化
func init() {
	// 全局初始化配置,其他地方直接引用viper,不需要再初始化
	public.InitConf()
	// 设置日志级别
	tools.BuildLogger(viper.GetString("LOG_LEVEL"))
	// 初始化数据库
	model.InitDB()
}

func main() {
	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
