package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	ExternalServices struct {
		ExchangeAPIKey       string `yaml:"exchange_api_key"`
		AnotherServiceAPIKey string `yaml:"another_service_api_key"`
	} `yaml:"external_services"`

	Server struct {
		Network      string `yaml:"network"`
		Address      string `yaml:"address"`
		IdleTimeout  int    `yaml:"idle_timeout"`
		ReadTimeout  int    `yaml:"read_timeout"`
		WriteTimeout int    `yaml:"write_timeout"`
	} `yaml:"server"`
}

func LoadConfiguration(path string) (cfg *Config) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return
}
