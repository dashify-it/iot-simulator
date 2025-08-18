package simulator

import (
	"errors"
	"flag"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var FilePaths = make(map[string]*string)

func ReadFlags(flagNames []string) {
	for i := 0; i < len(flagNames); i++ {
		FilePaths[flagNames[i]] = flag.String(flagNames[i], "", "")
	}
	flag.Usage = func() {
		printHelp()
	}

	flag.Parse()

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Print(`
iot-sim: IoT Device Simulator
--------------------------------
This tool simulates IoT devices sending data to either an MQTT broker or a webhook API.

USAGE:
  iot-sim -config <config.yaml> -specs <specs.yaml>

OPTIONS:
  -config string   Path to the configuration YAML file (defines simulator configs)
  -specs string    Path to the message specs YAML file (defines devices & messages)
  -h, --help       Show this help message

--------------------------------
CONFIG FILE FORMAT (config.yaml):
--------------------------------
send-mqtt: false             # true = send via MQTT, false = send via API
mqtt:
  mqtt-host: localhost
  mqtt-port: 1883
  mqtt-user: user
  mqtt-password: user

api:
  endpoint: http://localhost:3000/send-data/
  api-key-header-name: x-api-key
  api-key: sk_test_8f93b2a7c4d14f2a9e8d1c5b7a9e3f12

--------------------------------
SPECS FILE FORMAT (specs.yaml):
--------------------------------
messages:
  - title: msg_1
    device: device_a
    type: string
    options:
      - first_msg
      - second_msg
      - third_msg
    rate: once

  - title: msg_2
    device: device_b
    type: int
    rate: 2pm

  - title: msg_3
    device: device_b
    type: int
    rate: 2pm
    max: 100
    min: 0

  - title: msg_4
    device: device_c
    type: object
    body:
      - title: msg_4_1
        type: int
        max: 100
        min: 0
      - title: msg_4_2
        type: decimal
        max: 200
        min: 1
    rate: 10pm

--------------------------------
NOTES:
- "rate" supports "once", "Xpm" (per minute), or other frequency options.
- "type" can be string, int, decimal, or object (nested fields).
- When send-mqtt = false, messages will be sent to the API endpoint instead.

Examples:
  iot-sim -config ./config.yaml -specs ./specs.yaml
`)
}

func readYamlFile(filePath, errMsg string) ([]byte, error) {
	if len(filePath) != 0 {
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return []byte{}, err
		}
		return fileData, nil
	}
	return []byte{}, errors.New(errMsg)
}

func ParseConfigFile() (Config, error) {
	cfg := Config{}
	cfgData, err := readYamlFile(*FilePaths["config"], "the config file was not provided")
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(cfgData, &cfg)
	return cfg, err
}

func ParseSpecFile() (Specs, error) {

	specs := Specs{}
	specsData, err := readYamlFile(*FilePaths["spec"], "the spec file was not provided")
	if err != nil {
		return specs, err
	}
	err = yaml.Unmarshal(specsData, &specs)
	return specs, err
}
