package simulator

import (
	"fmt"
)

func ReadConfigs() Config {
	cfg, err := ParseConfigFile()

	if err != nil {
		fmt.Println(err.Error())
	}
	return cfg
}
