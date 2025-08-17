package simulator

import (
	"errors"
	"flag"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var FilePaths = make(map[string]*string)

func ReadFlags(flagNames []string) {
	for i := 0; i < len(flagNames); i++ {
		FilePaths[flagNames[i]] = flag.String(flagNames[i], "", "")
	}
	flag.Parse()
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
