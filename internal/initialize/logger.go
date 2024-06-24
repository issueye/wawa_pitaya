package initialize

import (
	"pitaya/internal/config"
	"pitaya/internal/global"
	"pitaya/pkg/logger"
)

func InitLogger() {
	cfg := &logger.Config{
		Path:       config.GetParam(config.CfgLogPath, "log").String(),
		MaxSize:    config.GetParam(config.CfgLogMaxSize, "100").Int(),
		MaxAge:     config.GetParam(config.CfgLogMaxAge, "30").Int(),
		MaxBackups: config.GetParam(config.CfgLogMaxBackups, "30").Int(),
		Compress:   config.GetParam(config.CfgLogCompress, "true").Bool(),
		Level:      config.GetParam(config.CfgLogLevel, "0").Int(),
		Mode:       logger.LOM_RELEASE,
	}

	global.Logger, global.Log = logger.InitLogger(cfg)
}
