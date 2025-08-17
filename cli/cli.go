package cli

import (
	"fmt"

	"github.com/dashify-it/iot-sim/simulator"
)

func StartCli() {
	fmt.Println("Starting iot-sim")
	simulator.ReadFlags([]string{"config", "spec"})
	cfg := simulator.ReadConfigs()
	specs := simulator.ReadSpecs()
	simulator.Simulate(cfg, specs)
}
