package initialize

import "pitaya/internal/config"

// 初始化
func Initialize() {
	// 初始化运行目录
	RuntimeDir()
	// 初始化参数
	config.InitConfig()
	// 初始化日志
	InitLogger()
	// 初始化数据库
	InitDB()
}
