package global

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log    *zap.Logger
	Logger *zap.SugaredLogger
	// 全局配置
	DB         *gorm.DB
	MQTTClient mqtt.Client
)
