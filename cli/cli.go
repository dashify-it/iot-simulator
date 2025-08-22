package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/dashify-it/iot-sim/logger"
	"github.com/dashify-it/iot-sim/simulator"
)

func StartCli() {
	// Banner style
	bannerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("212")).
		Bold(true).
		Underline(true).
		Padding(1, 0)

	fmt.Println(bannerStyle.Render("🚀 IoT Simulator..."))

	logger.Log.Debug("Reading flags...")
	simulator.ReadFlags([]string{"config", "spec"})

	// Logging startup info
	logger.Log.Info("IoT Simulator started")

	cfg := simulator.ReadConfigs()
	specs := simulator.ReadSpecs()

	logger.Log.Info("Configuration and specs loaded, starting simulation...")
	simulator.Simulate(cfg, specs)
}
