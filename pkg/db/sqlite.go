package db

import (
	"time"

	sqlite "github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func InitSqlite(path string, log *zap.SugaredLogger) *gorm.DB {
	newLogger := glogger.New(
		Writer{
			log:    log,
			BPrint: true,
		},
		glogger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  glogger.Info,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	l := newLogger.LogMode(glogger.Silent)
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: l,
	})

	if err != nil {
		log.Panicf("连接数据库异常: %v", err)
		panic(err)
	}

	db = db.Debug()

	log.Infof("初始化sqlite数据库完成! dsn: %s", path)

	return db
}
