package simulator

import (
	"github.com/dashify-it/iot-sim/logger"
)

func ReadConfigs() Config {
	cfg, err := ParseConfigFile()

	if err != nil {
		logger.Log.Error("error parsing config file: ", err.Error())
	}
	return cfg
}
