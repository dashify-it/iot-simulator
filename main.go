package main

import (
	"github.com/dashify-it/iot-sim/cli"
	"github.com/dashify-it/iot-sim/logger"
)

func main() {
	logger.InitLogger(false)
	defer logger.Sync()

	cli.StartCli()
}
