package simulator

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MqttClient mqtt.Client

func InitMqttClient(cfg Config) {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.Mqtt.Host, cfg.Mqtt.Port))
	opts.SetClientID("simulator-client")
	opts.SetUsername(cfg.Mqtt.User)
	opts.SetPassword(cfg.Mqtt.Password)
	MqttClient = mqtt.NewClient(opts)

	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connected to MQTT broker")
}

func SendMqttMessage(topic string, message interface{}) error {
	token := MqttClient.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	fmt.Println("Message published.")
	return nil
}
