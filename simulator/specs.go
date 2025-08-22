package simulator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dashify-it/iot-sim/logger"
)

func ReadSpecs() Specs {
	specs, err := ParseSpecFile()

	if err != nil {
		logger.Log.Error("error parsing specs: ", err.Error())
	}
	specs.ValidateSpecs()
	if err != nil {
		logger.Log.Error("error validating specs: ", err.Error())
	}
	return specs
}

func (specs *Specs) ValidateSpecs() error {
	for i := 0; i < len(specs.Messages); i++ {
		err := specs.Messages[i].validateMsg()
		if err != nil {
			return err
		}
	}
	return nil
}

func (message *Message) validateMsg() error {
	message.SetDefaults()
	switch message.Type {
	case string(STRING):
		if len(message.Options) == 0 {
			return fmt.Errorf("options is missing on message %s", message.Title)
		}
	case string(OBJECT):
		if len(message.Body) == 0 {
			return fmt.Errorf("body is missing on message %s", message.Title)
		}
		if message.Type == "object" {
			for j := 0; j < len(message.Body); j++ {
				err := message.Body[j].validateMsg()
				if err != nil {
					return err
				}
			}
		}
	case string(INTEGER):
	case string(DECIMAL):
		if message.Max == message.Min {
			return fmt.Errorf("max and min values need to be provided and not equal for this message %s", message.Title)
		}
		if message.Max < message.Min {
			return fmt.Errorf("max val needs to be greater than min for this message %s", message.Title)
		}
	}
	return nil
}

func ExtractRate(rate string) (int, MessageRate) {
	_, isPS := parseRateTypeAndReturnNumber(rate, PS)
	if isPS {
		// enforce once per second for starter
		return 1, PS
	}
	res, isPM := parseRateTypeAndReturnNumber(rate, PM)
	if isPM {
		return res, PM
	}
	res, isPH := parseRateTypeAndReturnNumber(rate, PH)
	if isPH {
		return res, PH
	}
	res, isPD := parseRateTypeAndReturnNumber(rate, PD)
	if isPD {
		return res, PD
	}
	return 1, ONCE
}

func parseRateTypeAndReturnNumber(rate string, rateType MessageRate) (int, bool) {
	isRateType := strings.Contains(rate, string(rateType))
	if isRateType {
		result := strings.Split(rate, string(rateType))
		if len(result) == 1 {
			// does not has number
			return 1, true
		}
		i, err := strconv.Atoi(result[0])
		if err != nil {
			return 1, true
		}
		return i, true
	}
	return 1, false
}
