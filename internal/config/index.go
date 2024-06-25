package config

const (
	CfgServerMode    = "SERVER-MODE"
	CfgLogMaxSize    = "LOG-MAX-SIZE"
	CfgLogMaxBackups = "LOG-MAX-BACKUPS"
	CfgLogMaxAge     = "LOG-MAX-AGE"
	CfgLogCompress   = "LOG-COMPRESS"
	CfgLogLevel      = "LOG-LEVEL"
	CfgLogPath       = "LOG-PATH"
)

func InitConfig() {
	SetParamExist(CfgServerMode, "debug", `服务运行模式， debug \ release`)

	SetParamExist(CfgLogPath, "runtime/logs", "日志存放路径")
	SetParamExist(CfgLogMaxSize, "100", "日志大小")
	SetParamExist(CfgLogMaxBackups, "30", "最大备份数")
	SetParamExist(CfgLogMaxAge, "30", "保存天数")
	SetParamExist(CfgLogCompress, "true", "是否压缩")
	SetParamExist(CfgLogLevel, "-1", "日志输出等级")
}
