package simulator

import (
	"encoding/json"
	"fmt"

	"github.com/dashify-it/iot-sim/logger"
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
		logger.Log.Error(fmt.Sprintf("could not connect to mqtt broker %s", token.Error()))
		panic(token.Error())
	}
	logger.Log.Info(fmt.Sprintf("Connected to MQTT broker %s:%d", cfg.Mqtt.Host, cfg.Mqtt.Port))
}

func SendMqttMessage(topic string, message interface{}) error {
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		logger.Log.Error("Error parsing json: ", err.Error())
		return err
	}
	token := MqttClient.Publish(topic, 0, false, fmt.Sprintf("%v", string(jsonBytes)))
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	logger.Log.Info(fmt.Sprintf("Message published successfully %s to topic: %s", fmt.Sprintf("%v", string(jsonBytes)), topic))
	return nil
}
