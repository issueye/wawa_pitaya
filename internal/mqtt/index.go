package mqtt

import (
	"fmt"
	"pitaya/internal/config"
	"pitaya/internal/global"
	"pitaya/pkg/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttClient() error {
	host := config.GetParam("mqtt_host", "").String()
	port := config.GetParam("mqtt_port", "").Int()

	if host == "" || port == 0 {
		global.Logger.Errorf("Mqtt host or port not set")
		return fmt.Errorf("Mqtt host or port not set")
	}

	// 如果已经连接，则断开连接，再重新连接
	if global.MQTTClient != nil {
		if global.MQTTClient.IsConnected() {
			global.MQTTClient.Disconnect(250)
			global.MQTTClient = nil
		}
	}

	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(fmt.Sprintf("pitaya_%d", utils.GenID()))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		global.Logger.Errorf("Connect mqtt failed: %v", token.Error())
		return token.Error()
	}

	global.MQTTClient = client

	err := SubProjects()
	if err != nil {
		global.Logger.Errorf("Subscribe project failed: %v", err)
		return err
	}

	return nil
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	global.Logger.Infof("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	global.Logger.Info("Connected")
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	global.Logger.Errorf("Connect lost: %v", err)
}
