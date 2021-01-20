package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ConfigYAML struct {
	Logging struct {
		LogFile     string `yaml:"logfile"`
		Level       string `yaml:"level"`
		LogToStdout bool   `yaml:"logToStdout"`
	}
	Defaults struct {
		Port int `yaml:"port"`
	}
	Bootstrap struct {
		Node struct {
			Address string `yaml:"address"`
			Port    int    `yaml:"port"`
		}
		Datacenter string `yaml:"datacenter"`
	}
}

func NewConfig(configFilePath string) (*ConfigYAML, error) {
	var config ConfigYAML
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Errorf("Unable to Open Config file at- %s, Error- %s", configFilePath, err)
		return nil, err
	}
	fmt.Printf(string(file))
	decoderError := yaml.Unmarshal(file, &config)
	if decoderError != nil {
		log.Errorf("Error while Parsing Config Yaml, Error- %s", decoderError)
		return nil, decoderError
	}
	log.Infof("Read YAML- \n%+v", config)
	return &config, nil
}

var Config *ConfigYAML

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
