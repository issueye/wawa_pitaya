package mqtt

import (
	"fmt"
	"pitaya/internal/global"
	"pitaya/internal/service"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SubProjects() error {
	srv := &service.ProjectList{}
	list, err := srv.List()
	if err != nil {
		global.Logger.Errorf("[SubProjects] error: %v", err)
		return err
	}

	// 订阅项目
	for _, project := range list {
		global.MQTTClient.Subscribe(project.Topic, 0, func(c mqtt.Client, m mqtt.Message) {
			// global.Logger.Infof("[SubProjects] %s: %s", m.Topic(), string(m.Payload()))
			fmt.Printf("[SubProjects] %s: %s\n", m.Topic(), string(m.Payload()))
			// 处理项目消息
			Sub(m.Topic(), m.Payload())
			// 将数据添加到列表中
		})
	}

	return nil
}

func Sub(topic string, data []byte) error {
	//

	return nil
}
