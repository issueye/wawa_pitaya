package initialize

import (
	"path/filepath"
	"pitaya/internal/global"
	"pitaya/internal/model"
	"pitaya/pkg/db"
)

func InitDB() {
	// 初始化数据库连接
	path := filepath.Join("runtime", "data", "data.db")
	global.DB = db.InitSqlite(path, global.Logger)

	InitData()
}

func InitData() {
	err := global.DB.Debug().AutoMigrate(
		&model.ProjectList{},
	)

	if err != nil {
		panic(err)
	}
}
