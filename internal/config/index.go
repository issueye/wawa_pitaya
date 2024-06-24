package config

const (
	CfgServerMode       = "SERVER-MODE"
	CfgServerPort       = "SERVER-PORT"
	CfgLogMaxSize       = "LOG-MAX-SIZE"
	CfgLogMaxBackups    = "LOG-MAX-BACKUPS"
	CfgLogMaxAge        = "LOG-MAX-AGE"
	CfgLogCompress      = "LOG-COMPRESS"
	CfgLogLevel         = "LOG-LEVEL"
	CfgLogPath          = "LOG-PATH"
	CfgServerApiName    = "API-NAME"
	CfgServerApiVersion = "api-version"
)

func InitConfig() {
	SetParamExist(CfgServerPort, "10065", "端口号")
	SetParamExist(CfgServerMode, "debug", `服务运行模式， debug \ release`)
	SetParamExist(CfgServerApiName, "api", "接口名称")
	SetParamExist(CfgServerApiVersion, "v1", "接口版本")

	SetParamExist(CfgLogPath, "log", "日志存放路径")
	SetParamExist(CfgLogMaxSize, "100", "日志大小")
	SetParamExist(CfgLogMaxBackups, "30", "最大备份数")
	SetParamExist(CfgLogMaxAge, "30", "保存天数")
	SetParamExist(CfgLogCompress, "true", "是否压缩")
	SetParamExist(CfgLogLevel, "-1", "日志输出等级")
}
