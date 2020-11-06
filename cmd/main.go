package main

import (
	"flag"
	"fmt"
	"github.com/iceberg98/go-consul-cleanup/pkg"
	"github.com/iceberg98/go-consul-cleanup/pkg/config"
	"github.com/iceberg98/go-consul-cleanup/pkg/logging"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "", "Path To the Configuration YAML file")
	flag.Parse()
	if err := config.ValidateConfigPath(configFilePath); err != nil {
		panic(err)
	}
	var err error
	config.Config, err = config.NewConfig(configFilePath)
	if err != nil {
		fmt.Printf("Error while reading Config file at Path- %s, Error- %s", configFilePath, err)
		panic(err)
	}
	logging.InitLogging()
	pkg.DeRegisterRedundantThings()
}
