package simulator

import (
	"fmt"
	"math/rand"
	"sync"
)

func Simulate(config Config, specs Specs) {
	if config.SendMqtt {
		InitMqttClient(config)
	}
	var wg sync.WaitGroup

	for i := range specs.Messages {
		msg := specs.Messages[i]
		msg.RateNumber, msg.RateType = ExtractRate(msg.Rate)
		wg.Add(1)
		go func(m Message) {
			defer wg.Done()
			HandleMessage(config, m)
		}(msg)
	}

	wg.Wait()
}
func HandleMessage(config Config, message Message) {
	body := buildMessage(message)
	var err error
	if config.SendMqtt {
		err = SendMqttMessage(message.Device, body)
	} else {
		err = SendApiRequest(config, body)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func buildMessage(message Message) interface{} {
	R := rand.New(rand.NewSource(55))
	body := map[string]interface{}{}
	if message.Type == string(STRING) {
		randomIndex := rand.Intn(len(message.Options))
		randomValue := message.Options[randomIndex]
		body[message.Title] = randomValue
	}
	if message.Type == string(INTEGER) {
		randomValue := R.Intn(int(message.Max)-int(message.Min+1)) + int(message.Min)
		body[message.Title] = randomValue
	}
	if message.Type == string(DECIMAL) {
		randomValue := message.Min + rand.Float64()*(message.Max-message.Min)
		body[message.Title] = randomValue
	}
	if message.Type == string(BOOLEAN) {
		randomValue := rand.Intn(2) == 1
		body[message.Title] = randomValue
	}
	if message.Type == string(OBJECT) {
		for i := range message.Body {
			body[message.Title] = buildMessage(message.Body[i])
		}
	}
	return body
}

func HandleMessageOld(sendMqtt bool, message Message) {
	medium := "mqtt"
	if !sendMqtt {
		medium = "api"
	}
	rate := "once"
	if message.RateType == PS {
		rate = fmt.Sprintf("%d per second", message.RateNumber)
	}
	if message.RateType == PM {
		rate = fmt.Sprintf("%d per minute", message.RateNumber)
	}
	if message.RateType == PH {
		rate = fmt.Sprintf("%d per hour", message.RateNumber)
	}
	if message.RateType == PD {
		rate = fmt.Sprintf("%d per day", message.RateNumber)
	}
	fmt.Printf("Handling %s msg from %s via %s for %s\n", message.Type, message.Device, medium, rate)
}
