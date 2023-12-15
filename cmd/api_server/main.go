package main

import (
	"flag"
	"log"
	"os"

	"github.com/O-Tempora/Ozon/config"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigPath = "config/default.yaml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to config file")
}

func main() {
	flag.Parse()
	config := &config.Config{}

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(bytes, config); err != nil {
		log.Fatal(err)
	}

}
