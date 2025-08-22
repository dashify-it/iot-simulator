package simulator

import (
	"math/rand"
	"sync"
	"time"

	"github.com/dashify-it/iot-sim/logger"
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
			runEvent(config, m)
		}(msg)
	}

	wg.Wait()
}
func runEvent(config Config, m Message) {
	switch m.RateType {
	case ONCE:
		// Just run once
		HandleMessage(config, m)

	case PS: // per second
		interval := time.Second / time.Duration(m.RateNumber)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			HandleMessage(config, m)
		}

	case PM: // per minute
		interval := time.Minute / time.Duration(m.RateNumber)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			HandleMessage(config, m)
		}

	case PH: // per hour
		interval := time.Hour / time.Duration(m.RateNumber)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			HandleMessage(config, m)
		}

	case PD: // per day
		interval := 24 * time.Hour / time.Duration(m.RateNumber)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			HandleMessage(config, m)
		}

	default:
		// fallback: just once if type is unknown
		HandleMessage(config, m)
	}
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
		logger.Log.Error("sending the msg failed: ", err.Error())
	}
}

func buildMessage(message Message) map[string]interface{} {
	R := rand.New(rand.NewSource(time.Now().UnixNano()))
	body := map[string]interface{}{}
	if message.Type == string(STRING) {
		randomIndex := R.Intn(len(message.Options))
		randomValue := message.Options[randomIndex]
		body[message.Title] = randomValue
	}
	if message.Type == string(INTEGER) {
		randomValue := R.Intn(int(message.Max)-int(message.Min+1)) + int(message.Min)
		body[message.Title] = randomValue
	}
	if message.Type == string(DECIMAL) {
		randomValue := message.Min + rand.Float64()*(message.Max-message.Min)
		body[message.Title] = float32(randomValue)
	}
	if message.Type == string(BOOLEAN) {
		randomValue := rand.Intn(2) == 1
		body[message.Title] = randomValue
	}
	if message.Type == string(OBJECT) {
		subMsg := map[string]interface{}{}
		for i := range message.Body {
			subMsg[message.Body[i].Title] = buildMessage(message.Body[i])[message.Body[i].Title]
		}
		body[message.Title] = subMsg
	}
	return body
}
