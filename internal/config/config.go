package config

import (
	"os"

	"github.com/DanielStefanK/twitchbot/internal/logger"
	"gopkg.in/yaml.v2"
)

var log = logger.NewLogger("config")

// Config parameter that are available
type Config struct {
	Bot struct {
		Username     string `yaml:"username"`
		OAuth        string `yaml:"oauth"`
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
		//MapsAPIToken string `yaml:"mapsAPIKey"`
		openWeatherAPI string `yaml:"openWeatherAPI"`
	} `yaml:"bot"`
}

// LoadConfig from the config file
func LoadConfig() *Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Error("could not read config file")
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Error("could not parse config file")
		panic(err)
	}

	return &cfg
}
