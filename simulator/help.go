package simulator

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrintHelp() {
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")). // cyan
		Bold(true)

	sectionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")). // purple
		Italic(true)

	fmt.Println(headerStyle.Render("IoT-Sim: IoT Device Simulator"))
	fmt.Println(headerStyle.Render("--------------------------------"))
	fmt.Println("This tool simulates IoT devices sending data to either an MQTT broker or a webhook API.")

	fmt.Println(sectionStyle.Render("USAGE:"))
	fmt.Println("  iot-sim -config <config.yaml> -specs <specs.yaml>")

	fmt.Println(sectionStyle.Render("OPTIONS:"))
	fmt.Println("  -config string   Path to the configuration YAML file (defines simulator configs)")
	fmt.Println("  -specs string    Path to the message specs YAML file (defines devices & messages)")
	fmt.Println("  -h, --help       Show this help message")

	fmt.Println(sectionStyle.Render("CONFIG FILE FORMAT (config.yaml):"))
	fmt.Println("--------------------------------")
	fmt.Println(`send-mqtt: false             # true = send via MQTT, false = send via API
mqtt:
  mqtt-host: localhost
  mqtt-port: 1883
  mqtt-user: user
  mqtt-password: user

api:
  endpoint: http://localhost:3000/send-data/
  api-key-header-name: x-api-key
  api-key: sk_test_8f93b2a7c4d14f2a9e8d1c5b7a9e3f12`)

	fmt.Println("\n" + sectionStyle.Render("SPECS FILE FORMAT (specs.yaml):"))
	fmt.Println("--------------------------------")
	fmt.Println(`messages:
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
    rate: 10pm`)

	fmt.Println("\n" + sectionStyle.Render("NOTES:"))
	fmt.Println(`- "rate" supports "once","Xps" (per second), "Xpm" (per minute), "Xph" (per hour) and "Xpd" (per day).
- "type" can be string, int, decimal, or object (nested fields).
- When send-mqtt = false, messages will be sent to the Webhook endpoint instead.`)

	fmt.Println("\n" + sectionStyle.Render("Examples:"))
	fmt.Println("  iot-sim -config ./config.yaml -specs ./specs.yaml")
}
